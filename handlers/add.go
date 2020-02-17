package handlers

import (
	"net/http"

	"github.com/invad0r/building-microservices-with-go/handlers/data"
)

// AddProduct ..
func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle Post Product.")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
