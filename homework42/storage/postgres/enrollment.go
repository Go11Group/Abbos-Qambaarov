package packages

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"mymod/module"

	_ "github.com/lib/pq"
)

type RepoNewEnrollment struct {
	Db *sql.DB
}



func (e *RepoNewEnrollment) CreateEnrollment(enrol modul.Enrollments) error {

    _, err := e.Db.Exec("insert into enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at) values ($1, $2, $3, $4, $5)",
        enrol.EnrollmentId, enrol.UserId, enrol.CourseId, time.Now(), time.Now())
    if err!= nil {
        return err
    }

    return nil
}



func (e *RepoNewEnrollment) GetAllEnrolment(filter modul.EnrollmentFilter) (*[]modul.Enrollments, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)

	query := `select enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at from enrollments  `

	fil := "where deleted_at = 0"

	if filter.Limit > 0 {
		params["limit"] = filter.Limit
		limit = ` LIMIT :limit`
	}

	if filter.Offset > 0 {
		params["limit"] = filter.Offset
		limit = ` LIMIT :offset`
	}


	query = query + fil + limit + offset

	query, arr = ReplaceQueryParamsEnrollment(query, params)
	rows, err := e.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	enrollments := []modul.Enrollments{}
	enrol := modul.Enrollments{}
	for rows.Next() {
		err = rows.Scan(&enrol.EnrollmentId, &enrol.UserId, &enrol.CourseId, &enrol.EnrollmentDate, &enrol.CreatedAt, &enrol.UpdatedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrol)
	}

	return &enrollments, nil
}

func ReplaceQueryParamsEnrollment(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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


func (e *RepoNewEnrollment) GetEnrollmentById(id string) (*modul.Enrollments, error) {
	
	rows, err := e.Db.Query("select * from enrollments where id = $1", id)
	if err != nil {
		return nil, err
	}

	enrol := modul.Enrollments{}

	err = rows.Scan(&enrol.EnrollmentId, enrol.UserId,&enrol.CourseId, enrol.EnrollmentDate, enrol.CreatedAt, enrol.UpdatedAt)
	if err!= nil {
        return nil, err
    }
	
	return &enrol, nil
}

func (e *RepoNewEnrollment) UpdateEnrollment(id string) error {
	
	_, err := e.Db.Exec("Update enrollments set updated_at = $1 where enrollment_id = $2",
		time.Now(),id)
	if err != nil {
		return err
	}

	return nil
}

func (e *RepoNewEnrollment)DeleteEnrollment(id string) error {
	
	_, err := e.Db.Exec("Delete from enrollment where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}