package dto

type ResultDto struct {
	Result interface{} `json:"result"`
	Meta   interface{} `json:"meta,omitempty"`
}

type PaginationMetaDto struct {
	Count int64 `json:"count"`
	Total int64 `json:"total"`
}
