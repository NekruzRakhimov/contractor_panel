package dto

import (
	"contractor_panel/domain/model"
	"net/url"
)

func ParseContractTemplateSearchParameters(values url.Values) (*model.ContractTemplateSearchParameters, error) {
	pagination, err := ParsePagination(values)
	if err != nil {
		return nil, err
	}

	return &model.ContractTemplateSearchParameters{
		Pagination: *pagination,
		Name:       ParseStringFilter(values, "name"),
	}, nil
}

type ContractTemplateDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func ConvertContractorToDto(template model.ContractTemplate) ContractTemplateDto {
	return ContractTemplateDto{Id: template.Id, Name: template.Name}
}

func ConvertContractTemplateDtos(list []model.ContractTemplate) []interface{} {
	result := make([]interface{}, len(list))

	for i := range list {
		result[i] = ConvertContractorToDto(list[i])
	}

	return result
}
