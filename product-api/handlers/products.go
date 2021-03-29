package handlers

import (
	"log"
	"net/http"
	"product-api/product-api/data"
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
		if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	//catch all
	w.WriteHeader(http.StatusMethodNotAllowed)

}
func (p*Products)addProduct(w http.ResponseWriter,r*http.Request)  {

}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	//data, err := json.Marshal(lp)
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}
	//w.Write(data)
}
