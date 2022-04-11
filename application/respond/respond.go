package respond

import (
	"contractor_panel/application/cerrors"
	"contractor_panel/application/dto"
	"contractor_panel/infrastructure/logging"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	contentTypeHeaderKey       = "Content-Type"
	contentTypeApplicationJson = "application/json"
	correlationIdHeaderKey     = "Correlation-ID"

	correlationIdCtxKey = "CorrelationId"
)

// AppErrorLog является вспомогательной структурой для логирования
// ошибки, так как невозможно экспортировать филды cerrors.AppError
type AppErrorLog struct {
	Error          string `json:"error"`
	HttpStatusCode int    `json:"httpStatusCode"`

	Code        int         `json:"code"`
	UserMessage string      `json:"userMessage"`
	Data        interface{} `json:"data"`
}

func With(w http.ResponseWriter, r *http.Request, result interface{}) {
	WithStatus(w, r, http.StatusOK, result, nil)
}

func WithMeta(w http.ResponseWriter, r *http.Request, result interface{}, meta interface{}) {
	WithStatus(w, r, http.StatusOK, result, meta)
}

func WithPagination(w http.ResponseWriter, r *http.Request, result []interface{}, total int64) {
	WithStatus(w, r, http.StatusOK, result, dto.PaginationMetaDto{Total: total, Count: int64(len(result))})
}

func WithStatus(w http.ResponseWriter, r *http.Request, statusCode int, result interface{}, meta interface{}) {
	rdto := &dto.ResultDto{Result: result, Meta: meta}

	cId := r.Context().Value(correlationIdCtxKey)
	if cId != nil {
		w.Header().Set(correlationIdHeaderKey, cId.(string))
	}
	w.Header().Set(contentTypeHeaderKey, contentTypeApplicationJson)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(rdto)
	if err != nil {
		panic("respond: " + err.Error())
	}
}

func WithError(w http.ResponseWriter, r *http.Request, err error) {
	cId := r.Context().Value(correlationIdCtxKey)
	if cId != nil {
		w.Header().Set(correlationIdHeaderKey, cId.(string))
	}
	w.Header().Set(contentTypeHeaderKey, contentTypeApplicationJson)

	var edto *dto.ErrorDto
	if ae, ok := err.(*cerrors.AppError); ok {
		edto = handleAppError(w, r, ae)
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		edto = handleValidationError(w, r, ve)
	} else {
		edto = handleGeneralError(w, r, err)
	}

	err = json.NewEncoder(w).Encode(edto)
	if err != nil {
		panic("respond: " + err.Error())
	}
}

func WithDto(w http.ResponseWriter, r *http.Request, statusCode int, dto interface{}) {

	cId := r.Context().Value(correlationIdCtxKey)
	if cId != nil {
		w.Header().Set(correlationIdHeaderKey, cId.(string))
	}
	w.Header().Set(contentTypeHeaderKey, contentTypeApplicationJson)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(dto)
	if err != nil {
		panic("respond: " + err.Error())
	}
}

func setCorrelationHeader(w http.ResponseWriter, r *http.Request) {
	cId := r.Context().Value(correlationIdCtxKey)
	if cId != nil {
		w.Header().Set(correlationIdHeaderKey, cId.(string))
	}
	w.Header().Set(contentTypeHeaderKey, contentTypeApplicationJson)
}

func handleAppError(w http.ResponseWriter, r *http.Request, err *cerrors.AppError) *dto.ErrorDto {
	logging.GetLogEntry(r).WithFields(logrus.Fields{"error": AppErrorLog{
		Error:          err.Error(),
		HttpStatusCode: err.HttpStatusCode(),
		Code:           err.Code(),
		UserMessage:    err.UserMessage(),
		Data:           err.Data(),
	}}).Error(err.UserMessage())

	w.WriteHeader(err.HttpStatusCode())
	return dto.NewErrorDto(err.Code(), err.UserMessage(), err.Data())
}

func handleValidationError(w http.ResponseWriter, r *http.Request, errs validator.ValidationErrors) *dto.ErrorDto {
	data := make([]map[string]interface{}, 0)
	for _, err := range errs {
		f := err.StructField()
		if err.StructNamespace() != "" {
			f = err.Namespace()
		}
		s := map[string]interface{}{
			"problem_param":   f,
			"problem_message": err.Error(),
		}

		data = append(data, s)
	}

	appErr := cerrors.NewAppError(
		errs, http.StatusBadRequest, cerrors.BadRequest, "возникли ошибки валидации тела запроса", data)
	logging.GetLogEntry(r).WithFields(logrus.Fields{"error": AppErrorLog{
		Error:          appErr.Error(),
		HttpStatusCode: appErr.HttpStatusCode(),
		Code:           appErr.Code(),
		UserMessage:    appErr.UserMessage(),
		Data:           data,
	}}).Error(appErr.UserMessage())

	w.WriteHeader(appErr.HttpStatusCode())
	return dto.NewErrorDto(appErr.Code(), appErr.UserMessage(), appErr.Data())
}

func handleGeneralError(w http.ResponseWriter, r *http.Request, err error) *dto.ErrorDto {
	logging.GetLogEntry(r).WithFields(logrus.Fields{"error": AppErrorLog{
		Error:          err.Error(),
		HttpStatusCode: http.StatusInternalServerError,
		Code:           cerrors.GeneralServiceError,
		UserMessage:    "",
		Data:           nil,
	}}).Error(err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	return dto.NewErrorDto(cerrors.GeneralServiceError, "необработанная ошибка сервиса", nil)
}
