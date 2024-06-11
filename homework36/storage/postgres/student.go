package packages

import (
	"database/sql"
	"strconv"
	// "os/user"
	modul "mymod/module"

	_ "github.com/lib/pq"
)

type RepoNewStudent struct {
	Db *sql.DB
}

func (u *RepoNewStudent) CreateStudent(student modul.Students) error {

	_, err := u.Db.Exec("insert into students(id, name, age, gender, course) values($1, $2, $3, $4, $5)",
				student.Id, student.Name, student.Age, student.Gender, student.Course)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewStudent) GetAllStudent(student modul.Students) (*[]modul.Students, error) {

	rows, err := u.Db.Query("select * from students")
	if err != nil {
		return nil, err
	}

	students := []modul.Students{}

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender, &student.Course)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func (u *RepoNewStudent) UpdateStudent(student modul.Students, id string) error {
	
	_, err := u.Db.Exec("Update students set course = $1 where id = $2", "Programming", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewStudent)DeleteStudent(student modul.Students, id string) error {
	
	Id, err := strconv.ParseInt(id,10,64)
	if err!= nil {
        return err
    }
	_, err = u.Db.Exec("Delete from students where id = $1", Id)
	if err != nil {
		return err
	}

	return nil
}
