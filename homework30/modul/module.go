package modul

type Users struct {
	Id       int
	UserName string
	Email    string
	Password string
}

type Products struct {
	Id            int
	Name          string
	Description   string
	Price         int
	StockQuantity int
}

type UsersProduct struct {
	Id       int
	UserName string
	Email    string
	Name     string
	Price    int
}

// --Users jadvali CREATE TABLE users ( id SERIAL PRIMARY KEY, username VARCHAR(50) NOT NULL,
// --email VARCHAR(100) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL );

// -- Products jadvali CREATE TABLE products ( id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL,
// --description TEXT, price NUMERIC(10, 2) NOT NULL, stock_quantity INT NOT NULL );
