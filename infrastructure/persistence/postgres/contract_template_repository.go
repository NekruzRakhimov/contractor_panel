package postgres

import (
	"context"
	"contractor_panel/domain/model"
	"github.com/jackc/pgx/v4/pgxpool"
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
