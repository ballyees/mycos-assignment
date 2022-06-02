package main

import (
	"database/sql"
	"fmt"

	"github.com/ballyees/mycos-assignment/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := InitialSQLiteDatabase("./db.db")
	panicOnError(err)
	defer db.Close() // close database connection when the program exits
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} [${time}]| ${status} - ${latency} ${method} | ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
	}))
	controllers.SetupRoutes(app, db)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	app.Listen(":8000")
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func InitialSQLiteDatabase(filename string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", filename)
	return
}
