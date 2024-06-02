package main

import (
	// "database/sql"
	// "fmt"
	"mymod/packages"
	// "os/user"

	_ "github.com/lib/pq"
)

func main()  {
	db, err := packages.Connectdb()
	if err != nil {
		panic(err)
	}

	// err = packages.CreateUser(db)
	// if err != nil {
	// 	panic(err)
	// }

	// GetAllUsers, err := packages.GetAllUser(db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _,v := range GetAllUsers {
	// 	fmt.Println(v)
	// }

	// err = packages.UpdateUser(db,1)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// err = packages.DeleteUser(db,2)		
	// if err != nil {
	// 	panic(err)
	// }





	err = packages.CreateProduct(db)
	if err != nil {
		panic(err)
	}

	// GetAllProducts, err := packages.GetAllProduct(db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _,v := range GetAllProducts {
	// 	fmt.Println(v)
	// }

	// err = packages.UpdateProduct(db,1)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// err = packages.DeleteProduct(db,2)		
	// if err != nil {
	// 	panic(err)
	// }



}