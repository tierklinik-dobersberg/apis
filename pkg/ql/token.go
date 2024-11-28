package ql

type Token int

type TokenInfo struct {
	Token   Token
	Literal string
}

var TokenLookup = map[Token]string{
	OTHER:              "OTHER",
	EOF:                "EOF",
	WS:                 "WS",
	STRING:             "STRING",
	EQUAL:              "=",
	GREATER_THAN:       ">",
	GREATER_THAN_EQUAL: ">=",
	LESS_THAN:          "<",
	LESS_THAN_EQUAL:    "<=",
	NOT_EQUAL:          "!=",
	AND:                "AND",
	OR:                 "OR",
	OPEN_BRACKET:       "(",
	CLOSE_BRACKET:      ")",
	PERCENT:            "%",
}

const (
	// Special Tokens
	OTHER Token = iota
	EOF
	WS

	// Main Literals
	STRING

	// Brackets
	OPEN_BRACKET
	CLOSE_BRACKET

	// Comparators
	GREATER_THAN
	GREATER_THAN_EQUAL
	LESS_THAN
	LESS_THAN_EQUAL
	EQUAL
	NOT_EQUAL
	PERCENT

	// Keywords
	AND
	OR
)
