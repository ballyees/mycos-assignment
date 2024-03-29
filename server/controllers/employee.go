package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/ballyees/mycos-assignment/database/model"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasttemplate"
)

const TimeFormat = "2006-01-02"

var tmpl = ReadFile("./controllers/template.html")
var rndr = fasttemplate.New(tmpl, "{{", "}}")

type EmployeePVD struct {
	Employee      model.Employee
	TotalPVD      float64
	TotalMonthPVD float64
}

func CalculatePVD(month float64, paidRate float64, employee model.Employee) float64 {
	return ((employee.Salary * paidRate / 100) * month) + ((employee.Salary * employee.PvdFundRate / 100) * month)
}

func CalculateTotalMonthPVD(employee model.Employee, now time.Time) (totalPVD float64) {
	t, _ := time.Parse(TimeFormat, employee.StartDate)
	diff := now.Sub(t)
	year := diff.Hours() / (24 * 365)
	month := math.Floor((year - math.Floor(year)) * 12)
	year = math.Floor(year)
	totalMonth := ((year * 12) + month) - 3
	if totalMonth <= 0 {
		return
	}
	if year > 5 { // over 5 year
		totalPVD = CalculatePVD(totalMonth, 80, employee)
	} else if year > 3 { // less than 5 year
		totalPVD = CalculatePVD(totalMonth, 50, employee)
	} else if year > 1 { // less than 3 year
		totalPVD = CalculatePVD(totalMonth, 30, employee)
	} else if year >= 0 { // less than 1 year
		totalPVD = CalculatePVD(totalMonth, 10, employee)
	}
	return
}

func CalculateTotalPVD(employee model.Employee, now time.Time) (totalPVD float64) {
	t, _ := time.Parse(TimeFormat, employee.StartDate)
	diff := now.Sub(t)
	year := diff.Hours() / (24 * 365)
	month := math.Floor((year - math.Floor(year)) * 12)
	year = math.Floor(year)
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
	totalPVD += CalculatePVD(math.Max((12*year)+month-3, 0), 10, employee) // less than 1 year
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
			employeePVDs[i].TotalMonthPVD = CalculateTotalMonthPVD(d, now)
		}
		// time.Sleep(time.Duration(10) * time.Second)
		return c.JSON(employeePVDs)
	}
}

func GetEmployeePVDHtml(db *sql.DB) func(c *fiber.Ctx) error {
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
			employeePVDs[i].TotalMonthPVD = CalculateTotalMonthPVD(d, now)
		}
		// time.Sleep(time.Duration(10) * time.Second)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		var reponse string

		for _, d := range employeePVDs {
			reponse += "<tr>\n"
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.ID)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.FirstName)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.LastName)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.BirthDate)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.StartDate)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.Salary)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.Employee.PvdFundRate)
			reponse += fmt.Sprintf("<td>%v</td>\n", d.TotalPVD)
			reponse += "</tr>\n"
		}
		_, err = rndr.Execute(c, map[string]interface{}{
			"tbody": reponse,
		})
		return err
	}
}

func ReadFile(fname string) (text string) {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		text = ""
		fmt.Printf("can't read file: %v\n", fname)
		log.Fatalln(err)
		return
	}

	// Convert []byte to string and print to screen
	text = string(content)
	return
}
