package database

import (
	"database/sql"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	_ "github.com/lib/pq"
)

func CreateConnection(driver, descriptor string, maxConns, maxIdle int) (connection *dbr.Connection, err error) {
	db, err := sql.Open(driver, descriptor)
	if err != nil {
		return
	}

	connection = &dbr.Connection{
		DB:            db,
		Dialect:       dialect.PostgreSQL,
		EventReceiver: &dbr.NullEventReceiver{},
	}

	connection.SetMaxOpenConns(maxConns)
	connection.SetMaxIdleConns(maxIdle)

	return
}

func SQLToString(stmt *dbr.SelectStmt) string {
	buf := dbr.NewBuffer()
	_ = stmt.Build(stmt.Dialect, buf)
	return buf.String()
}
