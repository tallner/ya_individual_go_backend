package services

import (
	"math/rand"
	"time"

	"cte.se/ya_individual_go_backend/data"
	"github.com/Pallinder/go-randomdata"
)

// var theRandom *rand.Rand
var theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

type PageView struct {
	Title                  string
	Rubrik                 string
	GITHUB_SHA_PLACEHOLDER string
}

type EmployeeService struct {
	// Any dependencies or data sources required by the service can be added here
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (s *EmployeeService) GetEmployees() ([]data.Employee, error) {
	var employees []data.Employee
	err := data.DB.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (s *EmployeeService) AddEmployee(employee *data.Employee) error {

	err := data.DB.Create(employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) AddManyEmployees() error {
	for i := 0; i < 10; i++ {
		employee := data.Employee{
			Age:  theRandom.Intn(50) + 18,
			Namn: randomdata.FirstName(randomdata.RandomGender),
			City: randomdata.City(),
		}
		err := data.DB.Create(&employee).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EmployeeService) GetEmployee(id string) (*data.Employee, error) {
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (s *EmployeeService) UpdateEmployee(id string, updatedEmployee *data.Employee) (*data.Employee, error) {
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if err != nil {
		return &employee, err
	}

	employee.Age = updatedEmployee.Age
	employee.Namn = updatedEmployee.Namn
	employee.City = updatedEmployee.City

	err = data.DB.Save(&employee).Error
	if err != nil {
		return &employee, err
	}
	return &employee, nil
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if err != nil {
		return err
	}

	err = data.DB.Delete(&employee).Error
	if err != nil {
		return err
	}
	return nil
}

// Add the GetPageView method to the EmployeeService
func (s *EmployeeService) GetPageView() *PageView {
	return &PageView{
		Title:                  "test",
		Rubrik:                 "Hello Golang",
		GITHUB_SHA_PLACEHOLDER: "Local",
	}
}
