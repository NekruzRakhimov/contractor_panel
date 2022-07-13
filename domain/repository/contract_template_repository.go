package repository

import (
	"context"
	"contractor_panel/domain/model"
)

type ContractTemplateRepository interface {
	FindContractTemplates(ctx context.Context,
		params model.ContractTemplateSearchParameters) ([]model.ContractTemplate, int64, error)
	GetContractTemplate(ctx context.Context, id int64) (*model.ContractTemplate, error)
	GetAllContracts(ctx context.Context, contractStatus string, userId int64) (contracts []model.ContractWithJsonB, err error)
	GetContractDetails(ctx context.Context, contractId int) (contract model.ContractWithJsonB, err error)
}
