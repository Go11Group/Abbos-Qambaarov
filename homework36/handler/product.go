package handler

import (
	"fmt"
	modul "mymod/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Handler2) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product := modul.Products{}

	products, err := p.PostProduct.GetAllProducts(product)
	if err != nil {
		panic(err)
	}
	for _,v := range *products {
		if id == v.Id {
			
			c.JSON(http.StatusOK,v)
			return

		}
	}
	c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found Product"})
	

}

func (p *Handler2) GetProduct(c *gin.Context) {

	product := modul.Products{}

	products, err := p.PostProduct.GetAllProducts(product)
	if err != nil {
		panic(err)
	}
	for _,v := range *products {
		c.JSON(http.StatusOK,v)
	}

}

func (p*Handler2) CreateProduct(c *gin.Context) {
	product := modul.Products{}

	err := c.BindJSON(&product)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error creating new product")
		return
	}

	err = p.PostProduct.CreateProducts(product)
	if err != nil {
		fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "Error creating new product")
        return
	}
	fmt.Println(product)
	c.JSON(http.StatusAccepted,"saved successfully")
}

func (p *Handler2) UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	newProduct := modul.Products{}

	err := c.BindJSON(&newProduct)
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest,"Error update product")
		return
	}
	product := modul.Products{}

	err = p.PostProduct.UpdateProducts(product,id)
	if err != nil {
		fmt.Println("Error update",err)
        c.JSON(http.StatusInternalServerError,"Error Update in database")
        return
	}
	c.JSON(http.StatusAccepted,"Saved successfully")
}

func (p *Handler2) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product := modul.Products{}
	err := p.PostProduct.DeleteProducts(product,id)
	if err!= nil {
        panic(err)
    }
	c.JSON(http.StatusAccepted,"Saved successfully")
}
