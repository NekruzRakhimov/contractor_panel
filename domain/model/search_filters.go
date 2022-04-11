package model

import (
	"math"
	"time"
)

type Pagination struct {
	Page          int64
	Size          int64
	ExternalTotal int64
}

func NewMaxPagination() *Pagination {
	return &Pagination{
		Page: 0,
		Size: math.MaxInt64,
	}
}

func (p Pagination) Limit() int64 {
	diff := p.Page*p.Size - p.ExternalTotal
	if diff < 0 {
		result := p.Size + diff
		if result < 0 {
			return 0
		}
		return result
	}

	return p.Size
}

func (p Pagination) Offset() int64 {
	result := p.Page*p.Size - p.ExternalTotal
	if result < 0 {
		return 0
	}
	return result
}

type DateFilter struct {
	From *time.Time
	To   *time.Time
}

func (d DateFilter) WithTime() *DateFilter {
	from := time.Date(d.From.Year(), d.From.Month(), d.From.Day(), 0, 0, 0, 0, d.From.Location())
	to := time.Date(d.To.Year(), d.To.Month(), d.To.Day(), 23, 59, 59, 999999999, d.To.Location())
	return NewDateFilter(&from, &to)
}

func NewDateFilter(from *time.Time, to *time.Time) *DateFilter {
	return &DateFilter{From: from, To: to}
}
