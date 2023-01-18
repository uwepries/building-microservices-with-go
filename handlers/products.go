package handlers

import (
	"log"
	"net/http"

	"github.com/uwepries/building-microservices-with-go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productsList := data.GetProducts()
	err := productsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
