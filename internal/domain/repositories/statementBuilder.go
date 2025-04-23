package repository

import (
	sq "github.com/Masterminds/squirrel"
)

var (
	StatementBuilder sq.StatementBuilderType
)

func InitialiseStatementBuilder() {
	StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
