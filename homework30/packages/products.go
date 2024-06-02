package packages

import (
	"database/sql"
	_ "github.com/lib/pq"
	"mymod/modul"
)

// type RepoNewUser struct{
// 	db *sql.DB
// }
func CreateProduct(db *sql.DB) error {
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer func()  {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	_, err = db.Exec("insert into product(id, names, description, price, stock_quantity, user_id) values($1, $2, $3, $4, $5, $6)",2, "oil", "good", 20000,18000,1)
	if err != nil {
		return err
	}

	return nil
}

func GetAllProduct(db *sql.DB) ([]modul.UsersProduct, error) {
	tr, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func()  {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	rows, err := db.Query("select u.id, u.username, u.email, p.names, p.price from users as u Left JOIN product as p ON u.id = p.user_id order by username")
	if err != nil {
		return nil, err
	}

	users := []modul.UsersProduct{}
	user := modul.UsersProduct{}

	for rows.Next() {
		err = rows.Scan(&user.Id,&user.UserName,&user.Email,&user.Name,user.Price)
		// if err != nil {
		// 	return nil, err
		// }
		users = append(users, user)
	}

	return users, nil
}


func UpdateProduct(db *sql.DB, id int) error {
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer func()  {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	_, err = db.Exec("Update product set price = $1 where id = $2",3000,id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	tr, err := db.Begin()
	if err != nil {
		return err
	}
	defer func()  {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	_, err = db.Exec("Delete from product where id = $1",id)
	if err != nil {
		return err
	}

	return nil
}