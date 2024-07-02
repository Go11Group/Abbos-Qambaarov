package postgres

import (
	"database/sql"
	models "mymod/module"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (s *UserRepo) CreateUser(user *models.Users) error {

	_, err := s.Db.Exec("insert into users(id, name, age, phone) values ($1, $2)",
		user.Id, user.Name, user.Age, user.Phone)

	return err
}

func (s *UserRepo) GetUserById(id string) (*models.Users, error) {
	var user = models.Users{Id: id}

	err := s.Db.QueryRow("select * from users where id = $1", id).
		Scan(&user.Name, &user.Age, &user.Phone)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepo) GetUser() (*[]models.Users, error) {
	rows, err := s.Db.Query("select * from users")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.Users
    for rows.Next() {
        var user = models.Users{}
        err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone)
        if err!= nil {
            return nil, err
        }
        users = append(users, user)
    }

    return &users, nil
}

func (s *UserRepo) UpdateUser(user models.Users ,id string) error {
	_, err := s.Db.Exec("update users set name = $1, age = $2, phone = $3 where id = $4",
        user.Name, user.Age, user.Phone, id)

    return err
}

func (s *UserRepo) DeleteUser(id string) error {
	_, err := s.Db.Exec("delete from users where id = $1", id)

    return err
}