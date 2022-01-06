package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)



func NewMock(logger log.Logger) (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		level.Error(logger).Log("error opening a stub database connection", err)
	}

	return db, mock
}

