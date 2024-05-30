package postgres

import (
	"database/sql"

	// "mymod/model"
)

// type CourseRepo struct {
// 	Db *sql.DB
// }


// func NewCourseRepo(db *sql.DB) *CourseRepo {
// 	return &CourseRepo{Db: db}
// }

// func (u *CourseRepo) GetAllCourses() ([]model.User, error) {
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



// func (u *CourseRepo) GetByCoursesID(id id) (model.User, error) {
// 	var user model.User

// 	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
// 					left join course c on c.id = s.course_id where s.id = $1`, id).
// 		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	return user, nil
// }

func CourseCreate(db *sql.DB) error {

	_, err := db.Exec("CREATE table course( ID serial primary key,	Name varchar, Field varchar )")
	if err != nil {
		panic(err)
	}
	return nil
}

func CourseInsert(db *sql.DB) error {
	_, err := db.Exec("insert into course(Name, Field) values($1,$2)","DizaynN23","Dizayner")
	if err != nil {
		panic(err)
	}
	return nil
}

func CourseUpdate(db *sql.DB) error {
	_, err := db.Exec("UPDATE course SET Name = $1 WHERE id = $2","DizaynN33",2)
	if err != nil {
		panic(err)
	}
	return nil
}

func CourseDelete(db *sql.DB) error {
	_, err := db.Exec("DELETE from course where id = $1",2)
	if err != nil {
		panic(err)
	}
	return nil
}

