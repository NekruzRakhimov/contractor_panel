package controller

import "github.com/gorilla/mux"

type RouteHandler interface {
	HandleRoutes(router *mux.Router)
}
