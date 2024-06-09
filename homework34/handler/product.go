package handler

import (
	"encoding/json"
	"fmt"
	modul "mymod/module"
	"net/http"
	"strings"
)

func (p *Handler2) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/getproductbyid/")
	product := modul.Products{}

	products, err := p.PostProduct.GetAllProducts(product)
	if err != nil {
		panic(err)
	}
	check := false
	for _,v := range *products {
		if id == v.Id {
			
			json.NewEncoder(w).Encode(v)
			check = true

		}
	}
	if check==false{

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	

}

func (p *Handler2) GetProduct(w http.ResponseWriter, r *http.Request) {
	product := modul.Products{}

	products, err := p.PostProduct.GetAllProducts(product)
	if err != nil {
		panic(err)
	}
	for _,v := range *products {
		json.NewEncoder(w).Encode(v)
	}
	if len(*products)==0{

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Empty database"))
		return
	}
	

}

func (p*Handler2) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := modul.Products{}

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}

	p.PostProduct.CreateProducts(product)

	fmt.Println(product)
	json.NewDecoder(r.Body).Decode(&product)
	w.Write([]byte("Successfully created"))
}

func (p *Handler2) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/updateProduct/")
	newProduct := modul.Products{}

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}
	product := modul.Products{}

	err = p.PostProduct.UpdateProducts(product,id)
	
		w.Write([]byte("Successfully updated"))
}

func (p *Handler2) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/deleteProduct/")

	product := modul.Products{}
	err := p.PostProduct.DeleteProducts(product,id)
	if err!= nil {
        panic(err)
    }
	
	w.Write([]byte("DELETE FROM DATABASE"))

}
