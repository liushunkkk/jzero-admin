package main

import (
	"server/tools/migrate/tables"
)

func main() {
	conn, err := tables.NewGormConn()
	if err != nil {
		panic(err)
	}
	err = tables.Migrate(conn)
	if err != nil {
		panic(err)
	}
}