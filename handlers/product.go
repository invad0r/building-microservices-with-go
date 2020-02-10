package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/invad0r/building-microservices-with-go/handlers/data"
)

// Products exported
type Products struct {
	l *log.Logger
}

// NewProducts exported
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts ..
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle GET Product.")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	//d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}

// AddProduct ..
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle Post Product.")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// UpdateProduct ..
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("handle PUT Product.", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found.", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found.", http.StatusInternalServerError)
		return
	}
}

// KeyProduct ..
type KeyProduct struct{}

// MiddlewareProductValidation ..
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	p.l.Println("running middleware..")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "unable to unmarshal json", http.StatusBadGateway)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})

}
