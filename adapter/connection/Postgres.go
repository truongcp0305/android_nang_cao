package connection

import "github.com/go-pg/pg/v10"

func Conn() *pg.DB {
	conn := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
	return conn
}
