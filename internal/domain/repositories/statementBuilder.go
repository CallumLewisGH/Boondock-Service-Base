// StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
package repository

import (
	"sync"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

// Database wraps the SQL connection pool
type StatementBuilder struct {
	StatementBuilder sq.StatementBuilderType
}

var (
	sbInstance *StatementBuilder
	once       sync.Once // Ensures initialization happens only once
)

func GetSB() *StatementBuilder {
	once.Do(func() {
		sbInstance = &StatementBuilder{}
		sbInstance.InitialiseStatementBuilder()
	})
	return sbInstance
}

func (sb *StatementBuilder) InitialiseStatementBuilder() {
	sb.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
