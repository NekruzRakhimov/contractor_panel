package dto

import (
	"contractor_panel/domain/model"
	"time"
)

type ContractorDto struct {
	Id        int64         `json:"id"`
	Resident  bool          `json:"resident"`
	Bin       *string       `json:"bin"`
	Name      *string       `json:"name"`
	Email     string        `json:"email" validate:"required"`
	BlockDate *time.Time    `json:"blockDate"`
	Status    string        `json:"status"`
	Employees []EmployeeDto `json:"employees"`
}

type EmployeeDto struct {
	Id        int64      `json:"id"`
	Email     string     `json:"email" validate:"required"`
	FullName  string     `json:"fullName"`
	Position  string     `json:"position" validate:"required"`
	BlockDate *time.Time `json:"blockDate"`
	Status    string     `json:"status"`
}

func ConvertContractor(c model.Contractor) ContractorDto {
	employees := make([]EmployeeDto, 0)

	for _, e := range c.Employees {
		employees = append(employees, ConvertContractorEmployee(e))
	}

	return ContractorDto{
		Id:        c.Id,
		Resident:  c.Resident,
		Bin:       c.Bin,
		Name:      c.Name,
		Email:     c.Email,
		BlockDate: c.BlockDate,
		Status:    string(c.Status),
		Employees: employees,
	}
}

func ConvertContractorEmployee(e model.Employee) EmployeeDto {
	return EmployeeDto{
		Id:        e.Id,
		Email:     e.Email,
		FullName:  e.FullName,
		Position:  e.Position,
		BlockDate: e.BlockDate,
		Status:    string(e.Status),
	}
}

