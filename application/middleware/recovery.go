package middleware

import (
	"contractor_panel/application/cerrors"
	"contractor_panel/application/respond"
	"contractor_panel/infrastructure/logging"
	"fmt"
	"net/http"
	"runtime/debug"
)

// RecoveryHandler за основу звят из gorilla/handlers

type recoveryHandler struct {
	handler    http.Handler
	printStack bool
}

// RecoveryOption provides a functional approach to define
// configuration for a handler; such as setting the logging
// whether or not to print stack traces on panic.
type RecoveryOption func(http.Handler)

func parseRecoveryOptions(h http.Handler, opts ...RecoveryOption) http.Handler {
	for _, option := range opts {
		option(h)
	}

	return h
}

// RecoveryHandler is HTTP middleware that recovers from a panic,
// logs the panic, writes http.StatusInternalServerError, and
// continues to the next handler.
//
// Example:
//
//  r := mux.NewRouter()
//  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//  	panic("Unexpected error!")
//  })
//
//  http.ListenAndServe(":1123", handlers.RecoveryHandler()(r))
func RecoveryHandler(opts ...RecoveryOption) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		r := &recoveryHandler{handler: h}
		return parseRecoveryOptions(r, opts...)
	}
}

// PrintRecoveryStack is a functional option to enable
// or disable printing stack traces on panic.
func PrintRecoveryStack(print bool) RecoveryOption {
	return func(h http.Handler) {
		r := h.(*recoveryHandler)
		r.printStack = print
	}
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			// Формируем кастомный ответ сервиса
			respond.WithError(w, r, cerrors.ErrInternalServerError(fmt.Errorf("%v", err)))
			h.log(r, err)
		}
	}()

	h.handler.ServeHTTP(w, r)
}

func (h recoveryHandler) log(r *http.Request, v ...interface{}) {
	entry := logging.GetLogEntry(r)
	entry.Error(v...)

	if h.printStack {
		stack := string(debug.Stack())
		entry.Errorln("stack:", stack)
	}
}
