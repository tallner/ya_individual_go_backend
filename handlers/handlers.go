package handlers

import (
	"net/http"

	"math/rand"

	"cte.se/ya_individual_go_backend/data"
	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)

type PageView struct {
	Title                  string
	Rubrik                 string
	GITHUB_SHA_PLACEHOLDER string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.html",
		&PageView{
			Title:                  "test",
			Rubrik:                 "Hello Golang",
			GITHUB_SHA_PLACEHOLDER: "Local"})
}

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func addEmployee(c *gin.Context) {

	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})

}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}

}

func getEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee data.Employee
	if err := data.DB.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func updateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee data.Employee
	if err := data.DB.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updatedEmployee data.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee data"})
		return
	}

	employee.Age = updatedEmployee.Age
	employee.Namn = updatedEmployee.Namn
	employee.City = updatedEmployee.City

	data.DB.Save(&employee)

	c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var employee data.Employee
	if err := data.DB.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	data.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.GET("/api/employees/:id", getEmployee) // Get a specific employee
	router.POST("/api/employees", addEmployee)
	router.PUT("/api/employees/:id", updateEmployee)    // Update an employee
	router.DELETE("/api/employees/:id", deleteEmployee) // Delete an employee

	return router
}
