package service

import (
	"context"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
	"encoding/json"
	"fmt"
	"log"
)

type ContractTemplateService interface {
	FindContractTemplates(ctx context.Context,
		params model.ContractTemplateSearchParameters) ([]model.ContractTemplate, int64, error)
	GetContractTemplate(ctx context.Context, id int64) (*model.ContractTemplate, error)
	DownloadContractTemplateFile(ctx context.Context, id int64) (string, error)
	GetAllContracts(ctx context.Context, contractType string, userId int64) (contractsMiniInfo []model.ContractMiniInfo, err error)
	GetContractDetails(ctx context.Context, contractID int) (contract model.Contract, err error)
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

func (s *contractTemplateService) GetAllContracts(ctx context.Context, contractType string, userId int64) (contractsMiniInfo []model.ContractMiniInfo, err error) {
	contractsWithJson, err := s.r.GetAllContracts(ctx, contractType, userId)
	if err != nil {
		return nil, err
	}

	contracts, err := ConvertContractsFromJsonB(contractsWithJson)
	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {
		contractMiniInfo := ConvertContractToContractMiniInfo(contract)
		contractsMiniInfo = append(contractsMiniInfo, contractMiniInfo)
	}

	return contractsMiniInfo, nil
}

func (s *contractTemplateService) GetContractDetails(ctx context.Context, contractId int) (contract model.Contract, err error) {
	fmt.Println("START GetContractDetails================")
	//contractWithJsonB, err := repository.GetContractDetails(contractId)
	contractWithJsonB, err := s.r.GetContractDetails(ctx, contractId)
	if err != nil {
		return model.Contract{}, err
	}

	contract, err = ConvertContractFromJsonB2(contractWithJsonB)
	if err != nil {
		return model.Contract{}, err
	}

	switch contractWithJsonB.Status {
	case "черновик":
		contract.Status = "DRAFT"
	case "на согласовании":
		contract.Status = "ON_APPROVAL"
	case "в работе":
		contract.Status = "ACTIVE"
	case "заверщённый":
		contract.Status = "EXPIRED"
	case "отменен":
		contract.Status = "CANCELED"
	default:
		contract.Status = "UNKNOWN"
	}

	return contract, nil

}

func ConvertContractToContractMiniInfo(contract model.Contract) (contractMiniInfo model.ContractMiniInfo) {
	if contract.Type == "marketing_services" {
		contractMiniInfo.ContractType = "Договор маркетинговых услуг"
	} else if contract.Type == "supply" {
		contractMiniInfo.ContractType = "Договор поставок"
	} else if contract.PrevContractId != 0 {
		contractMiniInfo.ContractType = "ДС"
	}

	switch contract.Status {
	case "черновик":
		contractMiniInfo.Status = "DRAFT"
	case "на согласовании":
		contractMiniInfo.Status = "ON_APPROVAL"
	case "в работе":
		contractMiniInfo.Status = "ACTIVE"
	case "заверщённый":
		contractMiniInfo.Status = "EXPIRED"
	case "отменен":
		contractMiniInfo.Status = "CANCELED"
	default:
		contractMiniInfo.Status = "UNKNOWN"
	}
	// здесь не получаю true

	contractMiniInfo.ID = contract.ID
	contractMiniInfo.ContractorName = contract.Requisites.ContractorName
	contractMiniInfo.ContractNumber = contract.ContractParameters.ContractNumber
	contractMiniInfo.Amount = contract.ContractParameters.ContractAmount
	contractMiniInfo.Author = contract.Manager
	contractMiniInfo.CreatedAt = contract.CreatedAt
	contractMiniInfo.UpdatedAt = contract.UpdatedAt
	contractMiniInfo.AdditionalAgreementNumber = contract.AdditionalAgreementNumber
	//contractMiniInfo.Status = contract.Status
	contractMiniInfo.Beneficiary = contract.Requisites.Beneficiary
	contractMiniInfo.IsExtendContract = contract.IsExtendContract
	contractMiniInfo.ExtendDate = contract.ExtendDate
	contractMiniInfo.StartDate = contract.ContractParameters.StartDate
	contractMiniInfo.EndDate = contract.ContractParameters.EndDate
	contractMiniInfo.ContractName = contract.ContractParameters.ContractName

	return contractMiniInfo
}

func ConvertContractsFromJsonB(contractsWithJsonB []model.ContractWithJsonB) (contracts []model.Contract, err error) {
	for _, contractWithJsonB := range contractsWithJsonB {
		contract, err := ConvertContractFromJsonB(contractWithJsonB)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func ConvertContractFromJsonB(contractWithJson model.ContractWithJsonB) (contract model.Contract, err error) {

	contract.ID = contractWithJson.ID
	contract.AdditionalAgreementNumber = contractWithJson.AdditionalAgreementNumber
	contract.Type = contractWithJson.Type
	//contract.Comment = contractWithJson.Comment
	contract.Manager = contractWithJson.Manager
	//contract.KAM = contractWithJson.KAM
	contract.Status = contractWithJson.Status
	//contract.CreatedAt = contractWithJson.CreatedAt
	//contract.UpdatedAt = contractWithJson.UpdatedAt
	//contract.WithTemperatureConditions = contractWithJson.WithTemperatureConditions
	//contract.PrevContractId = contractWithJson.PrevContractId
	//contract.IsExtendContract = contractWithJson.IsExtendContract
	//contract.ExtendDate = contractWithJson.ExtendDate
	//contract.DiscountBrand = contractWithJson.DiscountBrand
	//contract.ExtContractCode = contractWithJson.ExtContractCode

	err = json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)
	if err != nil {
		return model.Contract{}, err
	}

	//err = json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)
	//if err != nil {
	//	return model.Contract{}, err
	//}

	err = json.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)
	if err != nil {
		return model.Contract{}, err
	}

	//err = json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)
	//if err != nil {
	//	return model.Contract{}, err
	//}

	//err = json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)
	//if err != nil {
	//	return model.Contract{}, err
	//}

	//contract.IsExtendContract = contract.ContractParameters.IsExtendContract

	//contract.ExtendDate = contract.ContractParameters.ExtendDate
	return contract, nil
}

func ConvertContractFromJsonB2(contractWithJson model.ContractWithJsonB) (contract model.Contract, err error) {

	//log.Println("ConvertContractFromJsonB=======================", contractWithJson.ID, contractWithJson.IsExtendContract, contractWithJson.ExtendDate)
	contract.ID = contractWithJson.ID
	contract.AdditionalAgreementNumber = contractWithJson.AdditionalAgreementNumber
	contract.Type = contractWithJson.Type
	contract.Comment = contractWithJson.Comment
	contract.Manager = contractWithJson.Manager
	contract.KAM = contractWithJson.KAM
	contract.Status = contractWithJson.Status
	contract.CreatedAt = contractWithJson.CreatedAt
	contract.UpdatedAt = contractWithJson.UpdatedAt
	contract.WithTemperatureConditions = contractWithJson.WithTemperatureConditions
	contract.PrevContractId = contractWithJson.PrevContractId
	contract.IsExtendContract = contractWithJson.IsExtendContract
	contract.ExtendDate = contractWithJson.ExtendDate
	contract.DiscountBrand = contractWithJson.DiscountBrand
	contract.ExtContractCode = contractWithJson.ExtContractCode
	contract.View = contractWithJson.View

	//contract.Regions = contractWithJson.Regions

	err = json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)
	if err != nil {

		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Requisites), &contract.Requisites)] error is: ", err.Error())
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.SupplierCompanyManager), &contract.SupplierCompanyManager)] error is: ", err.Error())
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)
	if err != nil {
		log.Println("[service][.Unmarshal([]byte(contractWithJson.ContractParameters), &contract.ContractParameters)] error is: ", err.Error())
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Products), &contract.Products)] error is: ", err.Error())
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Discounts), &contract.Discounts)] error is: ", err.Error())
		return model.Contract{}, err
	}

	err = json.Unmarshal([]byte(contractWithJson.Regions), &contract.Regions)
	if err != nil {
		log.Println("[service][json.Unmarshal([]byte(contractWithJson.Regions), &contract.Regions)] error is: ", err.Error())
		return model.Contract{}, err
	}

	contract.IsExtendContract = contract.ContractParameters.IsExtendContract

	contract.ExtendDate = contract.ContractParameters.ExtendDate
	//log.Println("ДАННЫЕ ПО КОНТРАКТУ", contract)
	return contract, nil
}
