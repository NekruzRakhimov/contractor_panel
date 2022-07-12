package repository

import (
	"context"
	"contractor_panel/domain/model"
)

type ReportTemplateRepository interface {
	//FindContractTemplates(ctx context.Context)
	GetAllRBByContractorBIN(ctx context.Context, request model.RBRequest) (rbDTOs []model.RbDTO, err error)
	GetAllContractDetailByBIN(ctx context.Context, clientCode, PeriodFrom, PeriodTo string) (contracts []model.ContractWithJsonB, err error)
}
