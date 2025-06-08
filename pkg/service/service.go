package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/bufbuild/protovalidate-go"
	"github.com/sethvargo/go-envconfig"
	"github.com/tierklinik-dobersberg/apis/pkg/auth"
	"github.com/tierklinik-dobersberg/apis/pkg/cors"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/consuldiscover"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/wellknown"
	"github.com/tierklinik-dobersberg/apis/pkg/log"
	"github.com/tierklinik-dobersberg/apis/pkg/server"
	"github.com/tierklinik-dobersberg/apis/pkg/validator"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Mux struct {
	Public *http.ServeMux
	Admin  *http.ServeMux
	Shared *http.ServeMux
}

type Instance[Config any, Database any] struct {
	Name     wellknown.ServiceScope
	Config   Config
	Database Database
	Logger   *slog.Logger
	Catalog  discovery.Discoverer
	Clients  wellknown.Clients

	options []connect.Option

	Mux Mux

	publicServer *http.Server
	adminServer  *http.Server
	ctx          context.Context
}

func (i *Instance[Config, Database]) Context() context.Context {
	return i.ctx
}

func (i *Instance[Config, Database]) ConnectOptions() []connect.Option {
	return i.options
}

type DefaultConfigGetter interface {
	PublicAddress() string
	AdminAddress() string
	CORSConfig() *cors.Config
}

type BaseConfig struct {
	PublicListen   string   `env:"PUBLIC_LISTEN,default=:8080"`
	AdminListen    string   `env:"ADMIN_LISTEN,default=:8081"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS,default=*.dobersberg.vet"`
}

func (lc BaseConfig) PublicAddress() string { return lc.PublicListen }
func (lc BaseConfig) AdminAddress() string  { return lc.AdminListen }
func (lc BaseConfig) CORSConfig() *cors.Config {
	return &cors.Config{
		AllowedOrigins:   lc.AllowedOrigins,
		AllowCredentials: true,
	}
}

type ServiceConfig[Database any] interface {
	DatabaseCreator[Database]

	DefaultConfigGetter
}

func Configure[Database any, Config ServiceConfig[Database]](
	name wellknown.ServiceScope,

	// cfg is the service configuration
	cfg Config,

) (*Instance[Config, Database], error) {
	ctx := context.Background()

	// first, try to load the configuration
	if err := loadConfig(ctx, &cfg); err != nil {
		return nil, err
	}

	// setup the database if we got a database creator
	db, err := cfg.ConfigureDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to configure database: %w", err)
	}

	log := slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	catalog, err := consuldiscover.NewFromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to get service catalog client: %w", err)
	}

	// configure service clients
	clients := wellknown.ConfigureClients(wellknown.ConfigureClientOptions{
		Catalog: catalog,
	})

	sharedMux := http.NewServeMux()
	adminMux := http.NewServeMux()
	publicMux := http.NewServeMux()

	adminMux.Handle("/", sharedMux)
	publicMux.Handle("/", sharedMux)

	cors := cfg.CORSConfig()

	publicServer, err := server.CreateWithOptions(
		cfg.PublicAddress(),
		wrapWithKey("public", publicMux),
		server.WithCORS(*cors),
		server.WithBaseContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare public HTTP server: %w", err)
	}

	adminServer, err := server.CreateWithOptions(
		cfg.AdminAddress(),
		wrapWithKey("admin", publicMux),
		server.WithCORS(*cors),
		server.WithBaseContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare admin HTTP server: %w", err)
	}

	instance := &Instance[Config, Database]{
		Name:     name,
		Config:   cfg,
		Database: db,
		Logger:   log,
		Catalog:  catalog,
		Clients:  clients,
		Mux: Mux{
			Shared: sharedMux,
			Admin:  adminMux,
			Public: publicMux,
		},

		publicServer: publicServer,
		adminServer:  adminServer,
		ctx:          ctx,
	}

	instance.setupOptions()

	return instance, nil
}

func (i *Instance[Config, Database]) Run() error {
	// register the service at the catalog
	if err := discovery.Register(i.Context(), i.Catalog, &discovery.ServiceInstance{
		Name:    string(i.Name),
		Address: i.adminServer.Addr,
	}); err != nil {
		return fmt.Errorf("failed to rgister at service catalog: %w", err)
	}

	return server.Serve(i.ctx, i.publicServer, i.adminServer)
}

func (i *Instance[Config, Database]) setupOptions() {
	protoValidator, err := protovalidate.New()
	if err != nil {
		slog.Error("failed to prepare protovalidate", "error", err)
		os.Exit(1)
	}

	authInterceptor := auth.NewAuthAnnotationInterceptor(
		protoregistry.GlobalFiles,
		auth.NewIDMRoleResolver(i.Clients.RoleService),
		func(ctx context.Context, req connect.AnyRequest) (auth.RemoteUser, error) {
			serverKey, _ := ctx.Value(serverContextKey).(string)

			if serverKey == "admin" {
				return auth.RemoteUser{
					ID:          "service-account",
					DisplayName: req.Peer().Addr,
					RoleIDs:     []string{"idm_superuser"}, // FIXME(ppacher): use a dedicated manager role for this
					Admin:       true,
				}, nil
			}

			return auth.RemoteHeaderExtractor(ctx, req)
		},
	)

	interceptors := []connect.Interceptor{
		log.NewLoggingInterceptor(i.Logger),
		validator.NewInterceptor(protoValidator),
		authInterceptor,
	}

	i.options = []connect.Option{connect.WithInterceptors(interceptors...)}
}

var serverContextKey = struct{ S string }{S: "serverContextKey"}

func wrapWithKey(key string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), serverContextKey, key))

		next.ServeHTTP(w, r)
	})
}

func loadConfig(ctx context.Context, receiver any) error {
	if err := envconfig.Process(ctx, receiver); err != nil {
		return err
	}

	return nil
}
