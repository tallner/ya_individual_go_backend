package services

import (
	"math/rand"
	"testing"
	"time"

	"cte.se/ya_individual_go_backend/data"
	"github.com/stretchr/testify/assert"
)

func TestAddEmployee(t *testing.T) {
	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Add an employee
	err := service.AddEmployee()
	assert.NoError(t, err)
}

func TestAddManyEmployees(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Add many employees
	err := service.AddManyEmployees()
	assert.NoError(t, err)
}

func TestGetEmployees(t *testing.T) {
	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Get the employees
	employees, err := service.GetEmployees()
	assert.NoError(t, err)
	assert.NotNil(t, employees)
}

func TestGetEmployee(t *testing.T) {
	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Add an employee
	err := service.AddEmployee()
	assert.NoError(t, err)

	// Get the added employee
	employees, err := service.GetEmployees()
	assert.NoError(t, err)
	assert.NotNil(t, employees)

	// Get the first employee from the list
	employeeID := employees[0].ID
	employee, err := service.GetEmployee(employeeID)
	assert.NoError(t, err)
	assert.NotNil(t, employee)
}

func TestUpdateEmployee(t *testing.T) {
	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Add an employee
	err := service.AddEmployee()
	assert.NoError(t, err)

	// Get the added employee
	employees, err := service.GetEmployees()
	assert.NoError(t, err)
	assert.NotNil(t, employees)

	// Get the first employee from the list
	employeeID := employees[0].ID
	employee, err := service.GetEmployee(employeeID)
	assert.NoError(t, err)
	assert.NotNil(t, employee)

	// Update the employee
	updatedEmployee := &data.Employee{
		Age:  30,
		Namn: "Updated Name",
		City: "Updated City",
	}
	updatedEmployee, err = service.UpdateEmployee(employeeID, updatedEmployee)
	assert.NoError(t, err)
	assert.NotNil(t, updatedEmployee)
	assert.Equal(t, updatedEmployee.Age, 30)
	assert.Equal(t, updatedEmployee.Namn, "Updated Name")
	assert.Equal(t, updatedEmployee.City, "Updated City")
}

func TestDeleteEmployee(t *testing.T) {
	// Create a new instance of EmployeeService
	service := NewEmployeeService()

	// Add an employee
	err := service.AddEmployee()
	assert.NoError(t, err)

	// Get the added employee
	employees, err := service.GetEmployees()
	assert.NoError(t, err)
	assert.NotNil(t, employees)

	// Get the first employee from the list
	employeeID := employees[0].ID

	// Delete the employee
	err = service.DeleteEmployee(employeeID)
	assert.NoError(t, err)
}
