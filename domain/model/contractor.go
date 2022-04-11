package model

import (
	"fmt"
	"strconv"
	"time"
)

type Contractor struct {
	Id        int64
	Resident  bool
	Bin       *string
	Name      *string
	Email     string
	BlockDate *time.Time
	Status    ContractorStatus
	Employees []Employee
}

func (c Contractor) ReadModel(reader DbModelReader) (interface{}, error) {
	tmp := Contractor{}
	var employees []interface{}
	err := reader.Scan(&tmp.Id, &tmp.Resident, &tmp.Bin, &tmp.Name, &tmp.Email, &tmp.BlockDate, &tmp.Status, &employees)
	if err != nil {
		return nil, err
	}

	var employeeArray []Employee
	if len(employees) > 0 {
		for _, l := range employees {
			currentEmployee := l.(map[string]interface{})

			currentEmployeeIdStr := fmt.Sprintf("%.0f", currentEmployee["id"].(float64))
			currentContractorEmployeeIdStr := fmt.Sprintf("%.0f", currentEmployee["contractor_id"].(float64))
			currentEmployeeId, err := strconv.Atoi(currentEmployeeIdStr)
			if err != nil {
				return nil, err
			}
			currentContractorEmployeeId, err := strconv.Atoi(currentContractorEmployeeIdStr)
			if err != nil {
				return nil, err
			}

			var blockDate time.Time
			if currentEmployee["block_date"] != nil {
				blockDate, err = time.Parse(time.RFC3339, currentEmployee["block_date"].(string))
				if err != nil {
					return nil, err
				}
			}

			employeeArray = append(employeeArray, Employee{
				Id:           int64(currentEmployeeId),
				ContractorId: int64(currentContractorEmployeeId),
				Email:        currentEmployee["email"].(string),
				FullName:     currentEmployee["full_name"].(string),
				Position:     currentEmployee["position"].(string),
				BlockDate:    &blockDate,
				Status:       EmployeeStatus(currentEmployee["status"].(string)),
			})
		}
	}
	tmp.Employees = employeeArray

	return &tmp, nil
}

type Employee struct {
	Id           int64
	ContractorId int64
	Email        string
	FullName     string
	Position     string
	BlockDate    *time.Time
	Status       EmployeeStatus
}

type ContractorStatus string

const (
	ContractorStatusActive ContractorStatus = "ACTIVE"
	ContractorStatusBlock  ContractorStatus = "BLOCK"
)

type EmployeeStatus string

const (
	EmployeeStatusActive EmployeeStatus = "ACTIVE"
	EmployeeStatusBlock  EmployeeStatus = "BLOCK"
)

