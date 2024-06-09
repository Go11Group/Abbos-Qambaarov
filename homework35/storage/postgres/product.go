package packages

import (
	"database/sql"
	// "os/Products"
	modul "mymod/module"

	_ "github.com/lib/pq"
)

type RepoNewProducts struct {
	Db *sql.DB
}

func (u *RepoNewProducts) CreateProducts(product modul.Products) error {

	_, err := u.Db.Exec("insert into product(id, names, description, price, stock_quantity,user_id) values($1, $2, $3, $4, $5)", 
						product.Id, product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewProducts) GetAllProducts(product modul.Products) (*[]modul.Products, error) {

	rows, err := u.Db.Query("select * from product")
	if err != nil {
		return nil, err
	}

	products := []modul.Products{}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return &products, nil
}

func (u *RepoNewProducts) UpdateProducts(Products modul.Products, id string) error {
	
	_, err := u.Db.Exec("Update product set password = $1 where id = $2", "123456789", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewProducts)DeleteProducts(Products modul.Products, id string) error {
	
	_, err := u.Db.Exec("Delete from product where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
