package repository

import (
	"context"
	"contractor_panel/domain/model"
)

type ReportTemplateRepository interface {
	//FindContractTemplates(ctx context.Context)
	//GetAllRBByContractorBIN(ctx context.Context,userID int64 request model.RBRequest) (rbDTOs []model.RbDTO, err error)
	GetAllContractDetailByBIN(ctx context.Context, userID int64, request model.RBRequest) (contracts []model.ContractWithJsonB, err error)
}
