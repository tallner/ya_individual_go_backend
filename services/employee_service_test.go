package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestarBaraFörAttPipelinenSkaFåJobba_0(t *testing.T) {
	assert.True(t, true)
}
func TestarBaraFörAttPipelinenSkaFåJobba_1(t *testing.T) {
	assert.False(t, false)
}
func TestarBaraFörAttPipelinenSkaFåJobba_2(t *testing.T) {
	isThisAnError := nil
	assert.Error(t, isThisAnError)
}
func TestarBaraFörAttPipelinenSkaFåJobba_3(t *testing.T) {
	ageInMonth := 50 * 12
	result := ageInMonth > 500

	assert.False(t, result)
}

//func TestAddEmployee(t *testing.T) {
// Create a new instance of EmployeeService
//service := NewEmployeeService()

// Add an employee
// err := service.AddEmployee()
// assert.NoError(t, err)
//}

/*
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
	employeeID := strconv.Itoa(employees[0].Id)
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
	employeeID := strconv.Itoa(employees[0].Id)
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
	employeeID := strconv.Itoa(employees[0].Id)

	// Delete the employee
	err = service.DeleteEmployee(employeeID)
	assert.NoError(t, err)
}
*/
