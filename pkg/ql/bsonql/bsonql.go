package bsonql

import (
	"errors"

	"github.com/tierklinik-dobersberg/apis/pkg/ql"
	"go.mongodb.org/mongo-driver/bson"
)

type BSONQL struct {
	Schema ql.FieldProvider
}

func (b *BSONQL) Parse(s string) (bson.M, error) {
	res, err := ql.NewParser(s, b.Schema).Parse()
	if err != nil {
		return nil, err
	}

	return b.convertNode(res)
}

func (b *BSONQL) convertNode(node ql.Node) (bson.M, error) {
	switch v := node.(type) {
	case *ql.Expression:
		return b.convertExpression(v)
	case *ql.Operation:
		return b.convertOperation(v)
	}

	return nil, errors.New("unsupported node type")
}

func (b *BSONQL) convertExpression(exp *ql.Expression) (bson.M, error) {
	op := ""

	switch exp.Comparator {
	case "=":
		op = "$eq"
	case "!=":
		op = "$ne"
	case ">":
		op = "$gt"
	case ">=":
		op = "$gte"
	case "<":
		op = "$lt"
	case "<=":
		op = "$lte"
	}

	return bson.M{
		exp.Field.Name: bson.M{
			op: exp.Value,
		},
	}, nil
}

func (b *BSONQL) convertOperation(op *ql.Operation) (bson.M, error) {
	bsonOp := "$and"
	if op.Gate == "OR" {
		bsonOp = "$or"
	}

	left, err := b.convertNode(op.LeftNode)
	if err != nil {
		return nil, err
	}

	right, err := b.convertNode(op.RightNode)
	if err != nil {
		return nil, err
	}

	return bson.M{
		bsonOp: bson.A{
			left,
			right,
		},
	}, nil
}
