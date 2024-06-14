package packages

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"mymod/module"

	_ "github.com/lib/pq"
)

type RepoNewCourse struct {
	Db *sql.DB
}



func (c *RepoNewCourse) CreateCourse(course modul.Courses) error {

    _, err := c.Db.Exec("insert into courses (course_id, title, description, created_at) values ($1, $2, $3, $4)",
        course.CourseId,course.Title, course.Description, time.Now())
    if err!= nil {
		fmt.Println(err)
        return err
    }

    return nil
}



func (c *RepoNewCourse) GetAllCourse(filter modul.CourseFilter) (*[]modul.Courses, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)

	query := `select course_id, title, description, created_at, updated_at from courses  `

	fil := "where deleted_at = 0"

	if len(filter.Title) > 0 {
		params["name"] = filter.Title
		fil += " and gender = :title "
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

	query, arr = ReplaceQueryParamsCourse(query, params)
	rows, err := c.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	courses := []modul.Courses{}
	var course modul.Courses
	for rows.Next() {
		err = rows.Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return &courses,  nil
}

func ReplaceQueryParamsCourse(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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


func (c *RepoNewCourse) GetCourseById(id string) (*modul.Courses, error) {
	
	row := c.Db.QueryRow("select course_id, title, description, created_at, updated_at from users where id = $1", id)
	

	course := modul.Courses{} 

	err := row.Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	if err!= nil {
        return nil, err
    }
	
	return &course, nil
}

func (c *RepoNewCourse) UpdateCourse(course modul.Courses, id string) error {
	
	_, err := c.Db.Exec("Update courses set title = $1, description = $2, updated_at = $3 where course_id = $4",
		course.Title, course.Description ,time.Now(),id)
	if err != nil {
		return err
	}

	return nil
}

func (c *RepoNewCourse)DeleteCourse(id string) error {
	
	_, err := c.Db.Exec("update courses set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1 and deleted_at = 0", id)
	if err != nil {
		return err
	}

	return nil
}
