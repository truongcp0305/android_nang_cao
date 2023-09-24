package connection

import "github.com/go-pg/pg/v10"

func Conn() *pg.DB {
	conn := pg.Connect(&pg.Options{
		User:     "yenqt",
		Password: "o*!e2dev2RdevfffQ",
		Database: "sdocument_management_symper_vn",
		Addr:     "14.225.0.166:5432",
	})
	return conn
}
