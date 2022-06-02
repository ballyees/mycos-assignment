package controllers

import (
	"database/sql"
	"fmt"

	"github.com/ballyees/mycos-assignment/database/model"
	"github.com/ballyees/mycos-assignment/database/table"
	. "github.com/go-jet/jet/v2/sqlite"
)

func QueryAllEmployee(db *sql.DB) (error, []model.Employee) {
	stmt := SELECT(
		table.Employee.AllColumns,
	).FROM(
		table.Employee,
	)
	data := []model.Employee{}
	err := stmt.Query(db, &data)
	if err != nil {
		fmt.Println(err)
		return err, []model.Employee{}
	}
	return nil, data
}
