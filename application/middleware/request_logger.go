package middleware

import (
	"contractor_panel/infrastructure/logging"
	"fmt"
	"github.com/felixge/httpsnoop"
	"net/http"
	"net/url"
	"time"
)

// Код за основу звят из gorilla/handlers LoggingHandler

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	Request    *http.Request
	URL        url.URL
	TimeStamp  time.Time
	StatusCode int
	Size       int
}

// LogFormatter gives the signature of the formatter function passed to CustomLoggingHandler
type LogFormatter func(params LogFormatterParams)

// loggingHandler is the http.Handler implementation for LoggingHandlerTo and its
// friends

type loggingHandler struct {
	handler   http.Handler
	formatter LogFormatter
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	logger, w := makeLogger(w)
	rUrl := *req.URL

	h.handler.ServeHTTP(w, req)
	if req.MultipartForm != nil {
		_ = req.MultipartForm.RemoveAll()
	}

	params := LogFormatterParams{
		Request:    req,
		URL:        rUrl,
		TimeStamp:  t,
		StatusCode: logger.Status(),
		Size:       logger.Size(),
	}

	h.formatter(params)
}

func makeLogger(w http.ResponseWriter) (*responseLogger, http.ResponseWriter) {
	logger := &responseLogger{w: w, status: http.StatusOK}
	return logger, httpsnoop.Wrap(w, httpsnoop.Hooks{
		Write: func(httpsnoop.WriteFunc) httpsnoop.WriteFunc {
			return logger.Write
		},
		WriteHeader: func(httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
			return logger.WriteHeader
		},
	})
}

// writeRequestLog writes a log entry for req to w in Apache Combined Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
func writeRequestLog(params LogFormatterParams) {
	r := params.Request
	entry := logging.GetLogEntry(r)
	logFields := entry.Data

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	logFields["user_login"] = "-"
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method
	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	logFields["remote_addr"] = r.RemoteAddr
	logFields["referer"] = r.Referer()
	logFields["user_agent"] = r.UserAgent()

	logFields["resp_status_code"] = params.StatusCode
	logFields["resp_size"] = params.Size

	entry.WithFields(logFields).Info("request completed")
}

func RequestLoggerHandler(next http.Handler) http.Handler {
	return loggingHandler{next, writeRequestLog}
}
