package persistence

import (
	"contractor_panel/domain/model"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type AnonymousModelProvider func(dest ...interface{}) error

type DbScanner interface {
	// Read конвертирует первую найденную запись в модель и возвращает ссылку на нее.
	// При пустом наборе результатов возвращает nil.
	//
	// Пример использования:
	//
	//     tmp, _ := s.Read(SampleModel{})
	//     result := tmp.(*SampleModel)
	//
	Read(mp model.DbModelProvider) (interface{}, error)

	// ReadAll конвертирует набор результатов в массив моделей.
	//
	// Пример использования:
	//
	//     tmp, _ := s.ReadAll(SampleModel{})
	//     result := tmp.([]SampleModel)
	//
	ReadAll(mp model.DbModelProvider) (interface{}, error)

	// Scan считывает данные из первой найденной записи и возвращает true.
	// При пустом наборе результатов возвращает false.
	//
	// Пример использования:
	//
	//     var result string
	//     ok, _ := PgQuery(db, ctx, query).Scan(&result)
	//
	Scan(dest ...interface{}) (bool, error)
}

type DbRows interface {
	model.DbModelReader
	Close() error
	Next() bool
}

//region DbScannerWithError
type dbScannerWithError struct {
	err error
}

func NewDbScannerWithError(err error) DbScanner {
	return &dbScannerWithError{err: err}
}

func (s *dbScannerWithError) Read(model.DbModelProvider) (interface{}, error) {
	return nil, s.err
}

func (s *dbScannerWithError) ReadAll(model model.DbModelProvider) (interface{}, error) {
	return unwrapSlice(model, make([]interface{}, 0)), s.err
}

func (s *dbScannerWithError) Scan(...interface{}) (bool, error) {
	return false, s.err
}

//endregion

//region DefaultDbScanner
type defaultDbScanner struct {
	rows DbRows
}

func NewDbScanner(rows DbRows) *defaultDbScanner {
	return &defaultDbScanner{rows}
}

func (s *defaultDbScanner) Read(mp model.DbModelProvider) (interface{}, error) {
	defer s.rows.Close()

	for s.rows.Next() {
		l, err := mp.ReadModel(s.rows)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else if err != nil {
			return nil, err
		} else {
			return l, nil
		}
	}

	return nil, nil
}

func (s *defaultDbScanner) ReadAll(mp model.DbModelProvider) (interface{}, error) {
	defer s.rows.Close()

	result := make([]interface{}, 0)
	for s.rows.Next() {
		l, err := mp.ReadModel(s.rows)
		if errors.Is(err, sql.ErrNoRows) {
			return unwrapSlice(mp, make([]interface{}, 0)), nil
		} else if err != nil {
			return unwrapSlice(mp, make([]interface{}, 0)), err
		}

		result = append(result, l)
	}

	return unwrapSlice(mp, result), nil
}

func (s *defaultDbScanner) Scan(dest ...interface{}) (bool, error) {
	defer s.rows.Close()

	for s.rows.Next() {
		err := s.rows.Scan(dest...)
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}

	return false, nil
}

//endregion

// Преобразует запрос с именованными параметрами вида ":name" в запрос с позиционными параметрами
func InlineNamedPlaceholders(renamePlaceholder func(i uint8) string, query string, placeholders map[string]interface{}) (string, []interface{}, error) {
	var re = regexp.MustCompile(`(?m)[^:](:\w+)`)

	paramN := uint8(0)
	paramKeys := make(map[string]uint8)
	for _, match := range re.FindAllStringSubmatch(query, -1) {
		pKey := match[1]
		if _, ok := paramKeys[pKey]; !ok {
			paramKeys[pKey] = paramN
			paramN++
		}
	}

	inlineQuery := query
	inlinePlaceholders := make([]interface{}, paramN)
	for pKey, pN := range paramKeys {
		inlineQuery = strings.ReplaceAll(inlineQuery, pKey, renamePlaceholder(pN))
		if placeholderValue, ok := placeholders[pKey[1:]]; ok {
			inlinePlaceholders[pN] = placeholderValue
		} else {
			return "", nil, errors.New(fmt.Sprintf("не удалось найти значение параметра ':%s'", pKey))
		}
	}

	return inlineQuery, inlinePlaceholders, nil
}

// Преобразует массив типа []interface{} в массив []ModelType
func unwrapSlice(model interface{}, items []interface{}) interface{} {
	t := reflect.Indirect(reflect.ValueOf(model)).Type()
	arrT := reflect.SliceOf(t)
	arr := reflect.MakeSlice(arrT, len(items), len(items))
	for i := range items {
		v := reflect.Indirect(reflect.ValueOf(items[i]))
		arr.Index(i).Set(v)
	}

	return arr.Interface()
}
