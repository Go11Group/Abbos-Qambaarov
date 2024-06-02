package packages

import (
	"database/sql"
	// "os/user"
	"mymod/modul"

	_ "github.com/lib/pq"
)

// type RepoNewUser struct{
// 	db *sql.DB
// }
func CreateUser(db *sql.DB) error {
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
	_, err = db.Exec("insert into users(id, username, email, password) values($1, $2, $3, $4)",2, "Diyor", "diyor@gmain.com", "qwert")
	if err != nil {
		return err
	}

	return nil
}

func GetAllUser(db *sql.DB) ([]modul.UsersProduct, error) {
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

func UpdateUser(db *sql.DB, id int) error {
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
	_, err = db.Exec("Update users set password = $1 where id = $2","123456789",id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
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
	_, err = db.Exec("Delete from users where id = $1",id)
	if err != nil {
		return err
	}

	return nil
}

