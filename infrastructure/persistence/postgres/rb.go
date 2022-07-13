package postgres

import (
	"context"
	"contractor_panel/domain/model"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type ReportTemplateRepository struct {
	db *pgxpool.Pool
}

func NewReportTemplateRepository(db *pgxpool.Pool) *ReportTemplateRepository {
	return &ReportTemplateRepository{db: db}
}

//TODO: тут будет логика по БД, хотя

func (r *ReportTemplateRepository) GetAllRBByContractorBIN(ctx context.Context, request model.RBRequest) (rbDTOs []model.RbDTO, err error) {

	//TODO: необходимо, по ID брать бин хотя бы
	contractsWithJson, err := r.GetAllContractDetailByBIN(ctx, request.PeriodFrom, request.PeriodFrom)
	if err != nil {
		return nil, err
	}
	BulkConvertContractFromJsonB(contractsWithJson)

	return rbDTOs, nil
}
func BulkConvertContractFromJsonB(contractsWithJson []model.ContractWithJsonB) (contracts []model.Contract, err error) {
	for i := range contractsWithJson {
		contract, err := ConvertContractFromJsonB(contractsWithJson[i])
		if err != nil {
			log.Println("Error: service.BulkConvertContractFromJsonB. Error is: ", err.Error())
			continue
		}
		contracts = append(contracts, contract)
	}

	return
}

func ConvertContractFromJsonB(contractWithJson model.ContractWithJsonB) (contract model.Contract, err error) {
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

func (r ReportTemplateRepository) GetAllContractDetailByBIN(ctx context.Context, clientCode, PeriodFrom, PeriodTo string) (contracts []model.ContractWithJsonB, err error) {

	sqlQuery := `SELECT *
	         FROM (SELECT *
	         FROM "contracts"
	         WHERE (requisites ->> 'client_code' = ? AND
	         status in ('в работе', 'заверщённый'))) as sub_query
	         WHERE to_date(contract_parameters ->> 'start_date', 'DD.MM.YYYY') <= to_date(?, 'DD.MM.YYYY')
	         AND to_date(contract_parameters ->> 'end_date', 'DD.MM.YYYY') >= to_date(?, 'DD.MM.YYYY')`

	//if err = db.GetDBConn().Raw(sqlQuery, clientCode, PeriodFrom, PeriodTo).Scan(&contracts).Error; err != nil {
	//	log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
	//	return nil, err
	//}

	_, err = Query(r.db, ctx, sqlQuery, clientCode, PeriodFrom, PeriodTo).Scan(&contracts)
	if err != nil {
		log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
		return nil, err
	}

	//var brands []models.DiscountBrand
	//for i, contract := range contracts {
	//	if err = db.GetDBConn().Raw("SELECT id, brand as brand_name, brand_code, discount_percent FROM  brands  WHERE  contract_id = ?", contract.ID).Scan(&contracts[i].DiscountBrand).Error; err != nil {
	//		return nil, err
	//	}
	//
	//}

	return contracts, nil
}

//func GetAllContractDetailByBIN(ctx context.Context,clientCode, PeriodFrom, PeriodTo string, ) (contracts []model.ContractWithJsonB, err error) {
//	sqlQuery := `SELECT *
//          FROM (SELECT *
//          FROM "contracts"
//          WHERE (requisites ->> 'client_code' = ? AND
//          status in ('в работе', 'заверщённый'))) as sub_query
//          WHERE to_date(contract_parameters ->> 'start_date', 'DD.MM.YYYY') <= to_date(?, 'DD.MM.YYYY')
//          AND to_date(contract_parameters ->> 'end_date', 'DD.MM.YYYY') >= to_date(?, 'DD.MM.YYYY')`
//
//	Query(ctx)
//	if err = db.GetDBConn().Raw(sqlQuery, clientCode, PeriodFrom, PeriodTo).Scan(&contracts).Error; err != nil {
//		log.Println("[repository][GetAllContractDetailByBIN] error is: ", err.Error())
//		return nil, err
//	}
//
//	//var brands []models.DiscountBrand
//	for i, contract := range contracts {
//		if err = db.GetDBConn().Raw("SELECT id, brand as brand_name, brand_code, discount_percent FROM  brands  WHERE  contract_id = ?", contract.ID).Scan(&contracts[i].DiscountBrand).Error; err != nil {
//			return nil, err
//		}
//
//	}
//
//	return contracts, nil
//}
