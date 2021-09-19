package handler

import (
	"encoding/json"
	"net/http"
	"simple-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// getAll handles GET requests and returns all current products
func (handler *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	products, err := handler.Repository.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// getById handles GET requests and returns single product
func (handler *Handler) getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	product, err := handler.Repository.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// post handles POST requests to add new products
func (handler *Handler) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	product := data.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = handler.Repository.AddProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// put handles PUT requests to update products
func (handler *Handler) put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	values := data.Product{}
	json.NewDecoder(r.Body).Decode(&values)
	values.ID, _ = strconv.Atoi(vars["id"])

	err := handler.Repository.UpdateProduct(values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// delete handles DELETE requests and removes items from the database
func (handler *Handler) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := handler.Repository.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
