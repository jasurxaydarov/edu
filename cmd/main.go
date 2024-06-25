package main

import (
	"fmt"

	"github.com/jasurxaydarov/edu/api"
	"github.com/jasurxaydarov/edu/config"
	"github.com/jasurxaydarov/edu/pgx/db"
	"github.com/jasurxaydarov/edu/storage"
)

func main() {

	fmt.Println("start")
	cfg := config.Load()

	fmt.Println(cfg.PgConfig)
	pgxConn, err:=db.ConnectDB(cfg.PgConfig)

	if err != nil{

		return
	}

	fmt.Println(pgxConn)

	str:=storage.NewStorage(pgxConn)

	api.Api(str)

}
