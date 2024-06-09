package modul

type Users struct {
	Id       string
	UserName string
	Email    string
	Password string
}

type Products struct {
	Id            string
	Name          string
	Description   string
	Price         int
	StockQuantity int
}
type Students struct {
	Id     int
	Name   string
	Age    int
	Gender string
	Course string
}
