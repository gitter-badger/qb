package qb

import (
	"fmt"
	"strings"
)

const defaultDelimiter = "\n"

// NewQuery creates a new query and returns its pointer
func NewQuery() *Query {
	return &Query{
		clauses:      []string{},
		bindings:     []interface{}{},
		delimiter:    defaultDelimiter,
		bindingIndex: 0,
	}
}

// Query is the base abstraction for sql queries
type Query struct {
	clauses      []string
	bindings     []interface{}
	delimiter    string
	bindingIndex int
}

// SetDelimiter sets the delimiter of query
func (q *Query) SetDelimiter(delimiter string) {
	q.delimiter = delimiter
}

// AddClause appends a new clause to current query
func (q *Query) AddClause(clause string) {
	q.clauses = append(q.clauses, clause)
}

// AddBinding appends a new binding to current query
func (q *Query) AddBinding(bindings ...interface{}) {
	for _, v := range bindings {
		q.bindings = append(q.bindings, v)
	}
}

// Clauses returns all clauses of current query
func (q *Query) Clauses() []string {
	return q.clauses
}

// Bindings returns all bindings of current query
func (q *Query) Bindings() []interface{} {
	return q.bindings
}

// SQL returns the query struct sql statement
func (q *Query) SQL() string {

	if len(q.clauses) > 0 {
		sql := fmt.Sprintf("%s;", strings.Join(q.clauses, q.delimiter))
		return sql
	}

	return ""
}
