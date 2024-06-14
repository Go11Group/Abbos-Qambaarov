package packages

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	modul "mymod/module"

	_ "github.com/lib/pq"
)

type SearchFilter struct {
	Name  string 
	Email string
	Age1  int    
	Age2  int   
   }

type RepoNewUser struct {
	Db *sql.DB
}

func (u *RepoNewUser) CreateUser(user modul.Users) error {

	_, err := u.Db.Exec("insert into users(user_id,name, age, email, password, created_at) values($1, $2, $3, $4, $5, $6)",
		user.UserId, user.Name, user.Age, user.Email, user.Password, time.Now())
	if err != nil {
		return err
	}

	return nil
}


func (u *RepoNewUser) GetAllUser(filter modul.Filter) (*[]modul.Users, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)

	query := `select user_id, name, age,  email,  password,  created_at, updated_at from users  `

	fil := "where deleted_at = 0 "

	if len(filter.Name) > 0 {
		params["name"] = filter.Name
		fil += " and gender = :name "
	}

	if filter.Age > 0 {
		params["age"] = filter.Age
		fil += " and nation = :age "
	}

	if len(filter.Email) > 0 {
		params["email"] = filter.Email
		fil += " and age = :email "
	}

	if filter.Limit > 0 {
		params["limit"] = filter.Limit
		limit = ` LIMIT :limit`
	}

	if filter.Offset > 0 {
		params["limit"] = filter.Offset
		limit = ` LIMIT :offset`
	}

	query = query + fil + limit + offset

	query, arr = ReplaceQueryParamsUser(query, params)
	rows, err := u.Db.Query(query, arr...)

	users := []modul.Users{}
	var user modul.Users
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return &users, nil
}

func ReplaceQueryParamsUser(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)
	
	for k, v := range params {

		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func (u *RepoNewUser) GetUserById(id string) (*modul.Users, error) {
	user := modul.Users{}
	row:= u.Db.QueryRow("select user_id, name, age,  email,  password,  created_at, updated_at from users where user_id = $1", id)
	

	err := row.Scan(&user.UserId, &user.Name, &user.Age, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *RepoNewUser) UpdateUser(user modul.Users, id string) error {

	_, err := u.Db.Exec("Update users set name = $1, age = $2,  email = $3, password = $4, updated_at = $5 where user_id = $6",
		user.Name, user.Age, user.Email, user.Password, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (u *RepoNewUser) DeleteUser(id string) error {

	_, err := u.Db.Exec("update users set deleted_at = date_part('epoch', current_timestamp)::INT where user_id = $1 and deleted_at = 0", id)
	if err != nil {
		return err
	}

	return nil
}

func (us *RepoNewUser) GetUserCourses(id string) (*modul.UserCourses, error) {
	users := modul.UserCourses{UserId: id}
	rows, err := us.Db.Query(`
	SELECT c.course_id, c.title, c.description
	FROM courses c
	JOIN enrollments e ON c.course_id = e.course_id
	WHERE e.user_id = $1 AND c.deleted_at = 0
   `, id)
   
	for rows.Next() {
   
	 course := modul.Course2{}
	 err = rows.Scan(&course.CourseId, &course.Title, &course.Description)
	 if err != nil {
	  return nil, err
	 }
	 users.Courses = append(users.Courses, course)
	}
   
	return &users, err
   }
   
   func (us *RepoNewUser) SearchUsers(f SearchFilter) (*modul.SearchUser, error) {
	var (
	 params = make(map[string]interface{})
	 arr    []interface{}
	)
   
	query := `select user_id,
	   name,
	   email
	from users `
   
	filter := ``
   
	if len(f.Name) > 0 {
	 params["name"] = f.Name
	 filter += ` and name = :name `
	}
   
	if len(f.Email) > 0 {
	 params["email"] = f.Email
	 filter += ` and email = :email `
	}
   
	if f.Age1 > 0 && f.Age2 > 0 {
	 params["age1"] = f.Age1
	 params["age2"] = f.Age2
	 filter += ` AND EXTRACT(YEAR FROM age(birthday)) between :age1 and :age2 `
	}
   
	if len(filter) > 0 {
	 query = query + ` where deleted_at=0 ` + filter
	}
   
	query, arr = ReplaceQueryParamsUser(query, params)
	rows, err := us.Db.Query(query, arr...)
   
	if err != nil {
	 return nil, err
	}
   
	users := modul.SearchUser{}
	for rows.Next() {
   
	 user := modul.User2{}
   
	 err = rows.Scan(&user.UserID, &user.Name, &user.Email)
	 if err != nil {
	  return nil, err
	 }
	 users.Results = append(users.Results, user)
	}
   
	return &users, err
   }
   