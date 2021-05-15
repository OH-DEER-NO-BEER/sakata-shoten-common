package infrastructure

import (
	"database/sql"
	"sakata-shoten-common/interface/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func (s2 SqlHandler) Execute(s string, i ...interface{}) (database.Result, error) {
	panic("implement me")
}

func (s2 SqlHandler) Query(s string, i ...interface{}) (database.Row, error) {
	panic("implement me")
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
