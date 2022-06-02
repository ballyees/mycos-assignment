package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	v1 := app.Group("/v1")
	v1.Get("pvd", GetEmployeePVD(db))
}
