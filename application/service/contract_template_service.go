package service

import (
	"context"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
)

type ContractTemplateService interface {
	FindContractTemplates(ctx context.Context,
		params model.ContractTemplateSearchParameters) ([]model.ContractTemplate, int64, error)
	GetContractTemplate(ctx context.Context, id int64) (*model.ContractTemplate, error)
	DownloadContractTemplateFile(ctx context.Context, id int64) (string, error)
}

type contractTemplateService struct {
	r repository.ContractTemplateRepository
}

func NewContractTemplateService(r repository.ContractTemplateRepository) ContractTemplateService {
	return &contractTemplateService{r}
}

func (s *contractTemplateService) FindContractTemplates(ctx context.Context,
	params model.ContractTemplateSearchParameters) ([]model.ContractTemplate, int64, error) {
	return s.r.FindContractTemplates(ctx, params)
}

func (s *contractTemplateService) GetContractTemplate(ctx context.Context, id int64) (*model.ContractTemplate, error) {
	return s.r.GetContractTemplate(ctx, id)
}

func (s *contractTemplateService) DownloadContractTemplateFile(ctx context.Context, id int64) (string, error) {
	return "", nil
}
