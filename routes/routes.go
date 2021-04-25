// Package routes deals with the routing functionality
package routes

import (
	v1 "GithubSearch/handler/v1"
	"github.com/flannel-dev-lab/cyclops/v2/middleware"
	"github.com/flannel-dev-lab/cyclops/v2/router"
	"net/http"
)

// GetRoutes returns the registered routes for the handler
func GetRoutes(handler *v1.Handler) *router.Router {
	routerObj := router.New(false, nil, nil)

	routerObj.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	routerObj.Get("/v1/search", middleware.PanicHandler(middleware.AccessLogger(handler.Search)))

	return routerObj
}
