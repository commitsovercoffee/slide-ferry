package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/commitsovercoffee/slide-ferry/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {

		// expect the id in URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {

			p.l.Println("Invalid URI, more than one id.")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.l.Println("Invalid URI, more than one capture group.")
			p.l.Println(group[0])
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URI, unable to convert to number.")
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}

		p.l.Println("Got id : ", id)

		p.updateProducts(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle GET products")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal, please check json.", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found.", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found.", http.StatusInternalServerError)
		return
	}
}
