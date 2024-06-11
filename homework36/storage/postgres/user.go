package packages

import (
	"database/sql"
	// "os/user"
	modul "mymod/module"

	_ "github.com/lib/pq"
)

type RepoNewUser struct {
	Db *sql.DB
}

func (u *RepoNewUser) CreateUser(user modul.Users) error {

	_, err := u.Db.Exec("insert into users(id, username, email, password) values($1, $2, $3, $4)", user.Id, user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewUser) GetAllUser(user modul.Users) (*[]modul.Users, error) {

	rows, err := u.Db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	users := []modul.Users{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (u *RepoNewUser) UpdateUser(user modul.Users, id string) error {
	
	_, err := u.Db.Exec("Update users set password = $1 where id = $2", "123456789", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewUser)DeleteUser(user modul.Users, id string) error {
	
	_, err := u.Db.Exec("Delete from users where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
