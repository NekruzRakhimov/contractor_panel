package logging

import (
	"context"
	"contractor_panel/application/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

const (
	LogEntryCtxKey = "LogEntry"
)

var (
	defaultEntry *log.Entry
)

// ConfigureLogger конфигурирует общий логгер сервиса
func ConfigureLogger() {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "time",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
		PrettyPrint: viper.GetBool(config.LogPrettyPrint),
	})
	log.SetReportCaller(true)

	levelString := viper.GetString(config.LogLevel)
	level, err := log.ParseLevel(levelString)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		log.Warnf("invalid log level '%s', configured '%s'", levelString, log.GetLevel())
	} else {
		log.SetLevel(level)
	}

	log.AddHook(sequenceHook{})

	defaultEntry = log.WithFields(log.Fields{
		"app_name":     viper.GetString(config.AppName),
		"app_instance": viper.GetString(config.AppInstance),
	})
}

// GetLogEntry возвращает logrus.Entry, связанный с контекстом запроса r,
// иначе вернет формат лога по-умолчанию.
func GetLogEntry(r *http.Request) *log.Entry {
	if r == nil {
		return defaultEntry
	} else if entry, ok := r.Context().Value(LogEntryCtxKey).(*log.Entry); ok {
		return entry
	}

	return defaultEntry
}

// ContextWithLogEntry возвращает новый контекст, содержащий в себе запись логгера.
func ContextWithLogEntry(r *http.Request, entry *log.Entry) context.Context {
	return context.WithValue(r.Context(), LogEntryCtxKey, entry)
}
