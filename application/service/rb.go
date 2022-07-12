package service

import (
	"context"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
)

type ReportTemplateService interface {
	GetAllRBByContractorBIN(ctx context.Context, request model.RBRequest) (rbDTOs []model.RbDTO, err error)
}

type reportTemplateService struct {
	r repository.ReportTemplateRepository
}

func NewReportTemplateService(r repository.ReportTemplateRepository) *reportTemplateService {
	return &reportTemplateService{r: r}
}

func (s *reportTemplateService) GetAllRBByContractorBIN(ctx context.Context, request model.RBRequest) (rbDTOs []model.RbDTO, err error) {
	return s.r.GetAllRBByContractorBIN(ctx, request)

}
