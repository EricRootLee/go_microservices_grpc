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

func (p *Products) ServeHTTP(w http.ResponseWriter, response *http.Request) {
	lp := data.GetProducts()
	//data, err := json.Marshal(lp)
	 err :=  lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to parse data", http.StatusInternalServerError)
	}
	//w.Write(data)

}
