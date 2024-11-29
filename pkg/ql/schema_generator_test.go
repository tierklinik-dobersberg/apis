package ql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaGenerator(t *testing.T) {
	var model struct {
		ID     string
		Int    int
		Float  float64
		Nested struct {
			Field string
		}
	}

	res, err := SchemaFromModel(model, nil)
	require.NoError(t, err)

	require.Len(t, res, 4)
}
