package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"

	"log"
	"os"
)

var Db *gorp.DbMap

func init() {
	revel.OnAppStart(StartDb)
}

func StartDb() {
	db, _ := sql.Open("sqlite3", revel.BasePath+"/db")
	Db = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	Db.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	//TODO code here...

	err := Db.CreateTablesIfNotExists()
	if err != nil {
		revel.ERROR.Println(err.Error())
		panic(err)
	}
}
