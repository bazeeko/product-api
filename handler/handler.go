package handler

import (
	"net/http"
	"simple-api/data"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router     *mux.Router
	Repository data.ProductsRepository
}

type route struct {
	URL         string
	Method      string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

func NewHandler(repo *data.ProductsRepository) *Handler {
	return &Handler{
		Router:     mux.NewRouter(),
		Repository: *repo,
	}
}

func (handler *Handler) Init() {
	routes := []route{
		{
			URL:         "/products",
			Method:      "GET",
			HandlerFunc: handler.getAll,
		},
		{
			URL:         "/products/{id:[0-9]+}",
			Method:      "GET",
			HandlerFunc: handler.getById,
		},
		{
			URL:         "/products",
			Method:      "POST",
			HandlerFunc: handler.post,
		},
		{
			URL:         "/products/{id:[0-9]+}",
			Method:      "PUT",
			HandlerFunc: handler.put,
		},
		{
			URL:         "/products/{id:[0-9]+}",
			Method:      "DELETE",
			HandlerFunc: handler.delete,
		},
	}

	for _, route := range routes {
		handler.Router.HandleFunc(route.URL, route.HandlerFunc).Methods(route.Method)
	}

}
