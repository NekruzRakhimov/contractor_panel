package postgres

import (
	"contractor_panel/domain/model"
	"fmt"
	"github.com/lib/pq"
	"reflect"
	"strings"
)

func EscapeLikeFilterValue(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "%", "\\%")
	value = strings.ReplaceAll(value, "_", "\\_")

	return value
}

func AppendStringLikeFilter(filters *string, args model.NamedArguments, columnName string, value *string, pattern string) {
	if value == nil {
		return
	}
	tmpValue := strings.ToUpper(*value)

	tmpValue = EscapeLikeFilterValue(tmpValue)

	filterKey := genFilterKey(args)
	args[filterKey] = fmt.Sprintf(pattern, tmpValue)

	*filters = *filters + fmt.Sprintf(" and upper(%s) like :%s", columnName, filterKey)
}

func AppendDateFilter(filters *string, args model.NamedArguments, columnName string, value *model.DateFilter) {
	if value == nil {
		return
	}

	if value.From != nil && value.From == value.To {
		filterKey := genFilterKey(args)
		args[filterKey] = *value.From

		*filters = *filters + fmt.Sprintf(" and %s = :%s", columnName, filterKey)
		return
	}

	if value.From != nil {
		fromFilterKey := genFilterKey(args)
		args[fromFilterKey] = *value.From

		*filters = *filters + fmt.Sprintf(" and %s >= :%s", columnName, fromFilterKey)
	}

	if value.To != nil {
		toFilterKey := genFilterKey(args)
		args[toFilterKey] = *value.To

		*filters = *filters + fmt.Sprintf(" and %s <= :%s", columnName, toFilterKey)
	}
}

func AppendEqualsFilter(filters *string, args model.NamedArguments, columnName string, value interface{}) {
	if isNilPtrValue(value) {
		return
	}

	filterKey := genFilterKey(args)
	args[filterKey] = value

	*filters = *filters + fmt.Sprintf(" and %s=:%s", columnName, filterKey)
}

func AppendNotEqualsFilter(filters *string, args model.NamedArguments, columnName string, value interface{}) {
	if isNilPtrValue(value) {
		return
	}

	filterKey := genFilterKey(args)
	args[filterKey] = value

	*filters = *filters + fmt.Sprintf(" and (%s is null or %s<>:%s)", columnName, columnName, filterKey)
}

func AppendInListFilter(filters *string, args model.NamedArguments, columnName string, value interface{}) {
	if isNilListValue(value) {
		return
	}

	filterKey := genFilterKey(args)
	args[filterKey] = pq.Array(value)

	*filters = *filters + fmt.Sprintf(" and %s=any(:%s)", columnName, filterKey)
}

func AppendNotInListFilter(filters *string, args model.NamedArguments, columnName string, value interface{}) {
	if isNilListValue(value) {
		return
	}

	filterKey := genFilterKey(args)
	args[filterKey] = pq.Array(value)

	*filters = *filters + fmt.Sprintf(" and %s<>all(:%s)", columnName, filterKey)
}

func AppendPagination(filters *string, args model.NamedArguments, pagination model.Pagination) {
	args["limit"] = pagination.Limit()
	args["offset"] = pagination.Offset()

	*filters = *filters + " limit :limit offset :offset"
}

func genFilterKey(args model.NamedArguments) string {
	return fmt.Sprintf("filter%d", len(args))
}

func isNilListValue(value interface{}) bool {
	return value == nil || reflect.ValueOf(value).IsNil()
}

func isNilPtrValue(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
