package router

import "go-boiler-plate/handler/health"

func (r *router) HealthRoutes() {
	r.HandleFunc("/health", health.Health)
}
