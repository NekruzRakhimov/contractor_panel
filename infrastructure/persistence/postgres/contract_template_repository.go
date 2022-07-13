package postgres

import (
	"context"
	"contractor_panel/domain/model"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type ContractTemplateRepository struct {
	db *pgxpool.Pool
}

func NewContractTemplateRepository(db *pgxpool.Pool) *ContractTemplateRepository {
	return &ContractTemplateRepository{db}
}

func (r *ContractTemplateRepository) FindContractTemplates(ctx context.Context,
	params model.ContractTemplateSearchParameters) ([]model.ContractTemplate, int64, error) {
	args := model.NamedArguments{}
	args["id"] = model.ContractTemplateDictionaryCode
	queryTotal := `select count(*)`
	querySelect := `select v.id, v.value, v.file`
	queryFrom := ` from dictionaries d 
						left join dictionary_values v on v.dictionary_id = d.id`
	filters := ` where d.id = :id  and d.is_removed = false`

	AppendStringLikeFilter(&filters, args, "v.value", params.Name, "%s%%")

	var total int64
	_, err := QueryWithMap(r.db, ctx, queryTotal+queryFrom+filters, args).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return []model.ContractTemplate{}, 0, nil
	}

	paginatedFilters := filters + ` order by v.id desc`
	AppendPagination(&paginatedFilters, args, params.Pagination)

	result, err := QueryWithMap(r.db, ctx, querySelect+queryFrom+paginatedFilters, args).ReadAll(model.ContractTemplate{})
	if err != nil {
		return nil, 0, err
	}

	return result.([]model.ContractTemplate), total, nil

}

func (r *ContractTemplateRepository) GetContractTemplate(ctx context.Context, id int64) (*model.ContractTemplate, error) {
	args := model.NamedArguments{}
	args["id"] = model.ContractTemplateDictionaryCode
	args["templateId"] = id
	query := `select v.id, v.value, v.file
				from dictionaries d 
					left join dictionary_values v on v.dictionary_id = d.id
				where d.id = :id and d.is_removed = false and v.id = :templateId`

	res, err := QueryWithMap(r.db, ctx, query, args).Read(model.ContractTemplate{})
	if err != nil {
		return nil, err
	}

	return r.unwrapContractTemplateSlice(res), nil
}

func (r *ContractTemplateRepository) unwrapContractTemplateSlice(res interface{}) *model.ContractTemplate {
	if res == nil {
		return nil
	} else {
		return res.(*model.ContractTemplate)
	}
}

func (r *ContractTemplateRepository) GetAllContracts(ctx context.Context, contractStatus string, userId int64) ([]model.ContractWithJsonB, error) {
	var name string
	var supplier string
	r.db.QueryRow(ctx, "SELECT name FROM contractors_contractor WHERE  id = $1", userId).Scan(&name)
	r.db.QueryRow(ctx, "SELECT requisites ->> 'beneficiary' AS  supplier  FROM contracts WHERE id = $1", 686).Scan(&supplier)
	fmt.Println("NAME OF CONTRAGENT", name)
	fmt.Println("NAME OF КОНТРАГЕНТ", supplier)

	fmt.Println("GetALlContract Calling---------------------------")

	//var brands2 model.Brand
	//var contract []model.ContractWithJsonB

	var contractStatusRus = ""
	//sqlQuery := "SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true"

	sqlQuery := "SELECT id, contract_parameters, requisites, manager, type, status,  case when (requisites ->> 'beneficiary') IS NULL then '' else requisites ->> 'beneficiary' end AS suppler  FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true"
	//sqlQuery := "SELECT id,  created_at  FROM contracts"

	//sqlQuery := "SELECT  id, type, status, requisites, manager,  contract_parameters," +
	//	" created_at, updated_at, is_individ, additional_agreement_number, ext_contract_code FROM contracts"

	log.Println("STATUS", contractStatus)
	fmt.Println("STATUS", contractStatus)
	fmt.Println("sqlQuery", sqlQuery)

	if contractStatus != "" && contractStatus != "ACTIVE_AND_EXPIRED" {
		switch contractStatus {
		case "DRAFT":
			contractStatusRus = "черновик"
		case "ON_APPROVAL":
			contractStatusRus = "на согласовании"
		case "ACTIVE":
			contractStatusRus = "в работе"
		case "EXPIRED":
			contractStatusRus = "заверщённый"
		case "CANCELED":
			contractStatusRus = "отменен"
		}
		sqlQuery += fmt.Sprintf(" AND status = '%s'", contractStatusRus)
	}

	if contractStatus == "ACTIVE_AND_EXPIRED" {
		sqlQuery += fmt.Sprintf(" AND status in ('%s', '%s')", "в работе", "заверщённый")
	}

	//sqlQuery += " ORDER BY created_at DESC"
	sqlQuery += " ORDER BY id desc"

	//r.db.QueryRow(ctx, sqlQueryBrand).Scan(&brands2.ID, &brands2.Brand, &brands2.BrandCode, &brands2.DiscountPercent)
	//err = r.db.QueryRow(ctx, sqlQuery).Scan(&contract.ID, &contract.Status, &contract.Requisites)
	//if err != nil {
	//	fmt.Println("ERROR", err)
	//}
	//fmt.Println("Пример с контрактами", contract)

	//rows, err := r.db.Query(ctx, sqlQueryBrands)
	rows, err := r.db.Query(ctx, sqlQuery)
	//
	items := make([]model.ContractWithJsonB, 0)
	for rows.Next() {

		i := model.ContractWithJsonB{}

		//rows.Scan(&i.ID, &i.PrevContractId, &i.Status, &i.Requisites, &i.Manager, &i.Type, &i.SupplierCompanyManager, &i.ContractParameters, &i.Products, &i.Discounts, &i.Comment, &i.KAM, &i.UpdatedAt, &i.CreatedAt, &i.WithTemperatureConditions, &i.IsIndivid, &i.ExtContractCode)
		err := rows.Scan(&i.ID, &i.ContractParameters, &i.Requisites, &i.Manager, &i.Type, &i.Status, &supplier)
		fmt.Println("КОНТРАГЕНТ", supplier)
		//fmt.Println("контракты внутри цикла ", i)

		//err := rows.Scan(&i.ID, &i.Type, &i.Status, &i.Requisites, &i.Manager, &i.ContractParameters, &i.CreatedAt, &i.UpdatedAt, &i.IsIndivid, &i.AdditionalAgreementNumber, &i.ExtContractCode)
		if err != nil {
			fmt.Println("ERROR", err)
		}
		if supplier == name {
			items = append(items, i)

		}

	}
	//fmt.Println("ARRAY", items)

	//fmt.Println("После выхода loop", items)

	//query, err := r.db.Query(ctx, sqlQuery)
	//query.Scan(&contracts)

	//_, err = Query(r.db, ctx, sqlQuery).Scan(&contracts)
	//fmt.Println("contracts RESULT", co ntracts)
	//contractsSL := make([]model.ContractWithJsonB, 0)
	//
	//for query.Next() {
	//	item := model.ContractWithJsonB{}
	//	query.Scan(&item.ID, &item.IsExtendContract, &item.ExtContractCode, &item.Discounts, &item.Products, &item.UpdatedAt)
	//	contractsSL = append(contractsSL, item)
	//
	//}
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ContractTemplateRepository) GetContractDetails(ctx context.Context, contractId int) (contract model.ContractWithJsonB, err error) {

	err = r.db.QueryRow(ctx, "SELECT id, type, status, requisites, manager, kam, contract_parameters, with_temperature_conditions, products,"+
		"discounts, comment, created_at, updated_at, is_individ,  additional_agreement_number,view, regions FROM contracts WHERE id = $1", contractId).Scan(

		&contract.ID, &contract.Type, &contract.Status, &contract.Requisites, &contract.Manager, &contract.KAM, &contract.ContractParameters, &contract.WithTemperatureConditions,
		&contract.Products, &contract.Discounts, &contract.Comment, &contract.CreatedAt, &contract.UpdatedAt, &contract.IsIndivid, &contract.AdditionalAgreementNumber, &contract.View, &contract.Regionsa)
	fmt.Println("договор", contract)
	if err != nil {
		log.Println("ошибка при Селекте", err)
		return model.ContractWithJsonB{}, err
	}

	return contract, nil
}
