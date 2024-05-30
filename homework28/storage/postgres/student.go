package postgres

import (
	"database/sql"

	// "mymod/model"
)

// type StudentRepo struct {
// 	Db *sql.DB
// }


// func NewStudentRepo(db *sql.DB) *StudentRepo {
// 	return &StudentRepo{Db: db}
// }

// func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
// 	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from student s
// 					left join course c on c.id = s.course_id `)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []model.User
// 	var user model.User
// 	for rows.Next() {
// 		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }



// func (u *StudentRepo) GetByStudentID(id id) (model.User, error) {
// 	var user model.User

// 	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
// 					left join course c on c.id = s.course_id where s.id = $1`, id).
// 		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	return user, nil
// }

func StudentCreate(db *sql.DB) error {

	_, err := db.Exec("CREATE table student( ID serial primary key,	Name varchar, Age int, Gender varchar, Course varchar )")
	if err != nil {
		panic(err)
	}
	return nil
}

func StudentInsert(db *sql.DB) error {
	_, err := db.Exec("insert into student(Name, Age, Gender, Course) values($1,$2,$3,$4)","Diyor",19,"m","Dizayner")
	if err != nil {
		panic(err)
	}
	return nil
}

func StudentUpdate(db *sql.DB) error {
	_, err := db.Exec("UPDATE student SET Age = $1 WHERE id = $2",20,1)
	if err != nil {
		panic(err)
	}
	return nil
}

func StudentDelete(db *sql.DB) error {
	_, err := db.Exec("DELETE from student where id = $1",2)
	if err != nil {
		panic(err)
	}
	return nil
}


