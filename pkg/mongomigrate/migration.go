package mongomigrate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Migrate describes the interface that is used to migrate a database to a new version.
type Migrate interface {
	Run(ctx mongo.SessionContext, cli *mongo.Database) error
}

// MigrateFunc is a utility method that implements the Migrate interface.
type MigrateFunc func(mongo.SessionContext, *mongo.Database) error

// Run applies the migration by calling up
func (up MigrateFunc) Run(ctx mongo.SessionContext, cli *mongo.Database) error {
	return up(ctx, cli)
}

// Migration is a database migration for a given data base.
type Migration struct {
	// Version holds the version of the database after the migration is executed.
	Version int

	// Description holds an optional description.
	Description string

	// Database is the name of the database that should be migrated.
	Database string

	// Up holds the actual Migrate interface that performs the migration under
	// a write transaction.
	// All MongoDB operations executed during the migration MUST use the provided
	// mongo.SessionContext since it captures the write-transaction.
	Up Migrate
}

// MigrationRecords are stored in the version database and are used to identify the current
// version of the database and which Migration(s) need to be applied.
type MigrationRecord struct {
	Database    string    `bson:"database"`
	Version     int       `bson:"version"`
	MigrateTime time.Time `bson:"migrateTime"`
}

// Migrator support applying schema migrations for multiple MongoDB databases.
// A migrator should typically be created at the start of your application and it's
// Run() method should be used before any data of your MongoDB databases and collections
// is accessed.
type Migrator struct {
	// verCol holds the name of the collection inside the database that holds migration records.
	// if unset, it defaults to "database-version"
	verCol string

	// lock to guard access to the migrations map
	lock sync.Mutex

	// migrations holds all migrations index by the database name they apply to.
	migrations map[string][]Migration

	// migrationDatabase holds that mongo.Database that contains the version-information collection.
	// This might be the same database as your application uses but might also be a completely different
	// and separate database.
	//
	// Note that using a different mongoDB instance to store the migration database other than the one the
	// client uses is not supported for the time being.
	migrationDatabase *mongo.Database
}

// NewMigrator returns a new migrator that has it's version information stored in a collection named verCol
// in the database passed as db.
// verCol defaults to "database-version" if not set.
// Using different mongodb instances for the migration database and the databases to migrate is not supported
// for the time being.
func NewMigrator(db *mongo.Database, verCol string) *Migrator {
	if verCol == "" {
		verCol = "database-version"
	}

	return &Migrator{
		verCol:            verCol,
		migrationDatabase: db,
		migrations:        make(map[string][]Migration),
	}
}

// Register registers one or more collections at the migrator.
// This function SHOULD be called before a call to Migrator.Run.
// Though, if a migration is registered afterwards, no error is returned.
// This is subject to change.
func (m *Migrator) Register(migrations ...Migration) {
	m.lock.Lock()
	defer m.lock.Unlock()

	sortNeeded := make(map[string]struct{})
	for _, value := range migrations {
		sortNeeded[value.Database] = struct{}{}
		m.migrations[value.Database] = append(m.migrations[value.Database], value)
	}

	for key := range sortNeeded {
		slices.SortStableFunc(m.migrations[key], func(a Migration, b Migration) int {
			return a.Version - b.Version
		})
	}
}

// Run applies all migrations registered at the migrator. All migrations (independed of the migrated database)
// are executed within a single transaction.
// Any error returned from a Migration will rollback all changes and cause the function to exit.
// This is required since multiple databases might be related to each other so successful migrations
// in one database must also be rollbacked if any other migration fails.
//
// If you want a different behavior and not fail all database migrations at once you shloud create a
// new Migrator for each set of connected databases. Note that it is still possible to use the same
// migration database and version-information collection.
func (m *Migrator) Run(ctx context.Context) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	session, err := m.migrationDatabase.Client().StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		col := m.migrationDatabase.Collection(m.verCol)

		for database, migrations := range m.migrations {
			// first, get the latest session record
			results := col.FindOne(ctx, bson.M{
				"database": database,
			}, options.FindOne().SetSort(bson.D{
				{
					Key:   "version",
					Value: -1,
				},
			}))

			// bail out on errors other than ErrNoDocuments
			if err := results.Err(); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
				return nil, fmt.Errorf("failed to fetch latest migration record of database %q: %w", database, err)
			}

			// decode the last migration record.
			var last MigrationRecord
			if results.Err() == nil {
				if err := results.Decode(&last); err != nil {
					return nil, fmt.Errorf("failed to decode migration record of database %q: %w", database, err)
				}
			}

			// get a database client for the database we are going to migrate.
			db := session.Client().Database(database)

			// find and apply missing migrations
			for _, mi := range migrations {
				if last.Version < mi.Version {
					if err := mi.Up.Run(ctx, db); err != nil {
						return nil, fmt.Errorf("failed to run migration to version %d: %w", mi.Version, err)
					} else {
						slog.Info("successfully migrated database", "database", database, "pervious-version", last.Version, "next-version", mi.Version)
						last.Version = mi.Version
					}
				}
			}

			// finnally, create a new database record
			if _, err := col.InsertOne(ctx, &MigrationRecord{
				Database:    database,
				Version:     last.Version,
				MigrateTime: time.Now(),
			}); err != nil {
				return nil, fmt.Errorf("failed to create new migration record for database %q: %w", database, err)
			}
		}

		return nil, nil
	})

	return err
}
