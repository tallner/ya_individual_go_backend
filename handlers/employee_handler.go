package handlers

import (
	"net/http"

	"cte.se/ya_individual_go_backend/data"
	"cte.se/ya_individual_go_backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService *services.EmployeeService
}

func NewEmployeeHandler(employeeService *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	employees, err := h.employeeService.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) AddEmployee(c *gin.Context) {
	var employee data.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee data"})
		return
	}

	err := h.employeeService.AddEmployee(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee added"})
}

func (h *EmployeeHandler) AddManyEmployees(c *gin.Context) {
	err := h.employeeService.AddManyEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add employees"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employees added"})
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	employee, err := h.employeeService.GetEmployee(employeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var updatedEmployee data.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee data"})
		return
	}

	employee, err := h.employeeService.UpdateEmployee(employeeID, &updatedEmployee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	err := h.employeeService.DeleteEmployee(employeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}

func (h *EmployeeHandler) GetPageView(c *gin.Context) {
	pageView := h.employeeService.GetPageView()
	c.HTML(http.StatusOK, "home.html", pageView)
}

func SetupRouter(employeeHandler *EmployeeHandler) *gin.Engine {
	router := gin.Default()
	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1234"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	router.LoadHTMLGlob("templates/**")

	router.GET("/", employeeHandler.GetPageView)
	router.GET("/api/getallemployees", employeeHandler.GetEmployees)
	router.POST("/api/addemployee", employeeHandler.AddEmployee)
	router.GET("/api/addmanyrandomemployees", employeeHandler.AddManyEmployees)
	router.GET("/api/getoneemployees/:id", employeeHandler.GetEmployee)
	router.PUT("/api/updateemployee/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/api/deleteemployee/:id", employeeHandler.DeleteEmployee)

	return router
}
