package infrastructure

import (
	"database/sql"
	"sakata-shoten-common/interface/database/repository"
)

type SqlHandler struct {
	Conn *sql.DB
}

func (s2 SqlHandler) Execute(s string, i ...interface{}) (repository.Result, error) {
	panic("implement me")
}

func (s2 SqlHandler) Query(s string, i ...interface{}) (repository.Row, error) {
	panic("implement me")
}

func NewSqlHandler() repository.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
