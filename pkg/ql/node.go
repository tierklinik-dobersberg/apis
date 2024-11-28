package ql

import "fmt"

const (
	OPERATION  = "operation"
	EXPRESSION = "expression"
)

// Node represents a node in the AST after the expression is parsed.
type Node interface {
	Type() string
	String() string
}

// Expression represents something like x=y or x>=y
type Expression struct {
	Field      *FieldSpec
	Comparator string
	Value      any
	Literal    string
}

// Operation represents a Node (Operation or Expression) compared with another Node using either `AND` or `OR`.
type Operation struct {
	LeftNode  Node
	Gate      string
	RightNode Node
}

// Type returns the type for expression.
func (e Expression) Type() string { return EXPRESSION }

// Type returns the type for operation.
func (o Operation) Type() string { return OPERATION }

// String returns the string representation of expression.
func (e Expression) String() string {
	if s, ok := e.Value.(string); ok {
		return fmt.Sprintf("%v %v %q", e.Field.Name, e.Comparator, s)
	}

	return fmt.Sprintf("%v %v %v", e.Field.Name, e.Comparator, e.Value)
}

// String returns the string representation of operation.
func (o Operation) String() string {
	if o.Gate == "" {
		return fmt.Sprintf("(%v)", o.LeftNode)
	}
	return fmt.Sprintf("(%v %v %v)", o.LeftNode, o.Gate, o.RightNode)
}
