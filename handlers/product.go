package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// GetProducts ..
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	//d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}
