package db

import (
	"context"
	"database/sql"
	"fmt"

	"CURD/internals/user/schema"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func Connect(dsn string) *bun.DB{

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, pgdialect.New())
	fmt.Println("pinging db...")
	if err := conn.Ping(); err != nil {
		fmt.Println("Connection failed...")
	} else {
		fmt.Println("Connected succusfuly!")
	}

	inital()

	return db
}


func inital()  {

	ctx := context.Background()
	_,err1 := db.NewCreateTable().
			 Model((*schema.Wser)(nil)).
			 IfNotExists().
		 	Exec(ctx)
	if err1 != nil{
		fmt.Println("table createion failed..")
	}
}