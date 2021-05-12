package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"go-boiler-plate/factory"
)

func NewCustomRouter(f factory.Factory, l *logrus.Logger) *router {
	r := &router{mux.NewRouter()}
	r.registerRoutes(f, l)

	return r
}

type router struct {
	*mux.Router
}

func (r *router) registerRoutes(f factory.Factory, l *logrus.Logger) {
	r.HealthRoutes()
}
