package controllers

import (
	"database/sql"
	"math"
	"time"

	"github.com/ballyees/mycos-assignment/database/model"
	"github.com/gofiber/fiber/v2"
)

type EmployeePVD struct {
	Employee model.Employee
	TotalPVD float64
}

func getSubMonth(month, max int) int {
	if month > max {
		return month - max
	}
	return 0
}

func CalculatePVD(month float64, paidRate float64, employee model.Employee) float64 {
	return ((employee.Salary * paidRate / 100) * month) + ((employee.Salary * employee.PvdFundRate / 100) * month)
}

func CalculateTotalPVD(employee model.Employee, now time.Time) float64 {
	t, _ := time.Parse(TimeFormat, employee.StartDate)
	diff := now.Sub(t)
	year := diff.Hours() / (24 * 365)
	month := math.Floor((year - math.Floor(year)) * 12)
	year = math.Floor(year)
	var totalPVD float64
	if year >= 5 { // over 5 year
		totalPVD += CalculatePVD(12*(year-5)+month, 80, employee)
		year = 5
		month = 0
	}
	if year >= 3 { // less than 5 year
		totalPVD += CalculatePVD(12*(year-3)+month, 50, employee)
		year = 3
		month = 0
	}
	if year >= 1 { // less than 3 year
		totalPVD += CalculatePVD(12*(year-1)+month, 30, employee)
		year = 1
		month = 0
	}
	if year >= 0 { // less than 1 year
		totalPVD += CalculatePVD(math.Max((12*year)+month-3, 0), 10, employee)
	}
	return totalPVD
}

func GetEmployeePVD(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err, data := QueryAllEmployee(db)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		now := time.Now()
		employeePVDs := make([]EmployeePVD, len(data))
		for i, d := range data {
			employeePVDs[i].Employee = d
			employeePVDs[i].TotalPVD = CalculateTotalPVD(d, now)
		}
		return c.JSON(employeePVDs)
	}
}
