package ql

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Basic(t *testing.T) {
	query := "(completed = true AND dueTime<=now) OR dueTime >= startOfWeek"

	completedField := FieldSpec{
		Name:         "completed",
		TypeResolver: BooleanType(),
	}
	dueTimeField := FieldSpec{
		Name: "dueTime",
	}

	p := NewParser(query, FieldList{completedField, dueTimeField})

	res, err := p.Parse()
	require.NoError(t, err)

	expected := &Operation{
		LeftNode: &Expression{
			Field:      &completedField,
			Comparator: "=",
			Value:      true,
			Literal:    "true",
		},
		Gate: "AND",
		RightNode: &Expression{
			Field:      &dueTimeField,
			Comparator: "<=",
			Value:      "now",
			Literal:    "now",
		},
	}

	spew.Dump(res)

	require.Equal(t, res.(*Operation).LeftNode.(*Expression).Value, true)

	if !assert.EqualValues(t, expected, res) {
		spew.Dump(expected)
		spew.Dump(res)
	}
}
