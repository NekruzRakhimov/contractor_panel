package dto

import (
	"contractor_panel/domain/model"
	"errors"
	"net/url"
	"strconv"
	"time"
)

const defaultPageSize = 10

func ParsePagination(values url.Values) (*model.Pagination, error) {
	var page int64 = 0
	var size int64 = defaultPageSize
	var err error

	pageStr := values.Get("page")
	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	sizeStr := values.Get("size")
	if sizeStr == "" {
		sizeStr = values.Get("count")
	}
	if sizeStr != "" {
		size, err = strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &model.Pagination{
		Page: page,
		Size: size,
	}, nil
}

func parseDateFilterValue(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	fromValue, err := time.Parse("02.01.2006", value)
	if err != nil {
		return nil, errors.New("invalid date filter")
	}

	return &fromValue, nil
}

func ParseDateFilter(values url.Values, key string) (*model.DateFilter, error) {
	dates, ok := values[key]
	if !ok {
		return nil, nil
	}

	if len(dates) != 2 {
		return nil, errors.New("invalid date filter")
	}

	fromValue, err := parseDateFilterValue(dates[0])
	if err != nil {
		return nil, err
	}

	toValue, err := parseDateFilterValue(dates[1])
	if err != nil {
		return nil, err
	}

	return model.NewDateFilter(fromValue, toValue), nil
}

func ParseStringFilter(values url.Values, key string) *string {
	statusFilterStr := values.Get(key)
	if statusFilterStr == "" {
		return nil
	}

	return &statusFilterStr
}
