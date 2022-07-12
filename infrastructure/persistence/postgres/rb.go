package postgres

import (
	"context"
	"contractor_panel/domain/model"
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

	return rbDTOs, nil
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
