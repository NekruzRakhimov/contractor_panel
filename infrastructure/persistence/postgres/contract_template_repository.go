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

func (r *ContractTemplateRepository) GetAllContracts(ctx context.Context, contractStatus string) (contracts []model.ContractWithJsonB, err error) {
	fmt.Println("GetALlContract Calling---------------------------")
	var brands model.Brand
	var brandsAlL []model.Brand
	var brands2 model.Brand

	var contractStatusRus = ""
	sqlQuery := "SELECT * FROM contracts WHERE id not in (select prev_contract_id from contracts) AND is_active = true"
	sqlQueryBrand := "SELECT id, brand, brand_code, discount_percent FROM brands WHERE id = 31"
	sqlQueryBrands := "SELECT id, brand, brand_code, discount_percent FROM brands"

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

	Query(r.db, ctx, sqlQueryBrand).Scan(&brands.ID, &brands.Brand, &brands.BrandCode, &brands.DiscountPercent)
	fmt.Println("Первый запрос", brands)
	r.db.QueryRow(ctx, sqlQueryBrand).Scan(&brands2.ID, &brands2.Brand, &brands2.BrandCode, &brands2.DiscountPercent)
	fmt.Println("Второй пример", brands2)

	//rows, err := r.db.Query(ctx, sqlQueryBrands)
	rows, err := r.db.Query(ctx, sqlQuery)

	items := make([]model.ContractWithJsonB, 0)
	for rows.Next() {
		i := model.ContractWithJsonB{}
		rows.Scan(&i.ID, &i.PrevContractId, &i.Status, &i.Requisites, &i.Manager, &i.Type, &i.SupplierCompanyManager, &i.ContractParameters, &i.Products, &i.Discounts, &i.Comment, &i.KAM, &i.UpdatedAt, &i.CreatedAt, &i.WithTemperatureConditions, &i.IsIndivid, &i.ExtContractCode)
		fmt.Println("внутри цикла ", i)
		items = append(items, i)

	}

	fmt.Println("После выхода loop", items)

	Query(r.db, ctx, sqlQueryBrands)

	fmt.Println("ВСЕ БРЕНДЫ", brandsAlL)

	query, err := r.db.Query(ctx, sqlQuery)
	//query.Scan(&contracts)

	//_, err = Query(r.db, ctx, sqlQuery).Scan(&contracts)
	fmt.Println("contracts RESULT", contracts)
	contractsSL := make([]model.ContractWithJsonB, 0)

	for query.Next() {
		item := model.ContractWithJsonB{}
		query.Scan(&item.ID, &item.IsExtendContract, &item.ExtContractCode, &item.Discounts, &item.Products, &item.UpdatedAt)
		contractsSL = append(contractsSL, item)

	}
	if err != nil {
		return nil, err
	}

	return contracts, nil
}
