package model

type ContractTemplate struct {
	Id   int64
	Name string
	Path string
}

const (
	ContractTemplateDictionaryCode = 9
)

func (c ContractTemplate) ReadModel(reader DbModelReader) (interface{}, error) {
	tmp := ContractTemplate{}
	err := reader.Scan(&tmp.Id, &tmp.Name, &tmp.Path)
	if err != nil {
		return nil, err
	}

	return &tmp, nil
}

type ContractTemplateSearchParameters struct {
	Pagination Pagination
	Name       *string
}
