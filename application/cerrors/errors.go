package cerrors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AppError struct {
	error          error
	httpStatusCode int

	code        int
	userMessage string
	data        interface{}
}

func NewAppError(err error, httpStatusCode int, code int, message string, data interface{}) *AppError {
	return &AppError{
		error:          err,
		httpStatusCode: httpStatusCode,
		code:           code,
		userMessage:    message,
		data:           data,
	}
}

// Возвращает текст ошибки
func (e *AppError) Error() string {
	if e.error == nil {
		return e.userMessage
	} else {
		return fmt.Sprintf("%s: %s", e.userMessage, e.error.Error())
	}
}

func (e *AppError) Unwrap() error {
	return e.error
}

func (e *AppError) HttpStatusCode() int {
	return e.httpStatusCode
}

func (e *AppError) Code() int {
	return e.code
}

func (e *AppError) UserMessage() string {
	return e.userMessage
}

func (e *AppError) Data() interface{} {
	return e.data
}

// region Коды ошибок

const (
	GeneralServiceError   = 50000
	BadRequest            = 50001
	ConfigurationError    = 50002
	ResourceNotFoundError = 50004

	CouldNotOpenDbConnection = 51000
	CouldNotPingDb           = 51001

	CouldNotSignIn      = 53000
	CouldNotVerifyToken = 53001

	CouldNotGetContractorById = 52000
)

// endregion

// region Ошибки

func ErrInternalServerError(err error) *AppError {
	return &AppError{
		error:          err,
		httpStatusCode: http.StatusInternalServerError,
		code:           GeneralServiceError,
		userMessage:    "необработанная ошибка сервиса",
	}
}

func ErrBadRequestVar(err error, name string) *AppError {
	ae := &AppError{}
	ae.error = err
	ae.httpStatusCode = http.StatusBadRequest
	ae.code = BadRequest
	ae.userMessage = fmt.Sprintf("ошибка валидации параметра `%s`", name)

	data := make([]map[string]interface{}, 0)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, err := range ve {
			s := map[string]interface{}{
				"problem_param":   name,
				"problem_message": fmt.Sprintf("значение не соответсвует тегу `%s`", err.Tag()),
			}

			data = append(data, s)
		}
	}

	ae.data = data

	return ae
}

func ErrCouldNotDecodeBody(err error) *AppError {
	return &AppError{
		error:          err,
		httpStatusCode: http.StatusBadRequest,
		code:           BadRequest,
		userMessage:    "возникла ошибка во время обработки тела запроса",
	}
}

func ErrConfigurationError(failedKeys []string) *AppError {
	return &AppError{
		httpStatusCode: http.StatusInternalServerError,
		code:           ConfigurationError,
		userMessage:    fmt.Sprintf("ошибка конфигурации сервиса. Не заданы следующие переменные окружения: %v", failedKeys),
		data:           failedKeys,
	}
}

func ErrResourceNotFound(r *http.Request) error {
	data := make(map[string]interface{}, 0)
	data["url"] = r.URL
	data["url_params"] = r.URL.Query()

	return &AppError{
		httpStatusCode: http.StatusNotFound,
		code:           ResourceNotFoundError,
		userMessage:    "ресурс по такому URL адресу не найден",
		data:           data,
	}
}

func ErrCouldNotConnectToDb(err error) *AppError {
	return &AppError{
		error:       err,
		code:        CouldNotOpenDbConnection,
		userMessage: "не удалось подключиться к базе данных",
	}
}

func ErrCouldNotPingDb(err error) *AppError {
	return &AppError{
		error:       err,
		code:        CouldNotPingDb,
		userMessage: "не удалось выполнить Ping комманду, возможно база данных неактивна",
	}
}

func ErrCouldNotGetContractorById(err error, id int64) *AppError {
	return &AppError{
		error:       err,
		code:        CouldNotGetContractorById,
		userMessage: fmt.Sprintf("ошибка во время запроса контрагента по ИД %d", id),
	}
}

func ErrCouldNotSignIn(err error, login string) *AppError {
	return &AppError{
		httpStatusCode: http.StatusUnauthorized,
		error:          err,
		code:           CouldNotSignIn,
		userMessage:    fmt.Sprintf("ошибка аутентификации пользователя по логину %s", login),
	}
}

func ErrCouldNotVerifyToken(err error) *AppError {
	return &AppError{
		httpStatusCode: http.StatusUnauthorized,
		error:          err,
		code:           CouldNotVerifyToken,
		userMessage:    "ошибка авторизации",
	}
}

// endregion
