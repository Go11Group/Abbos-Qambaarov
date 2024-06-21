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

type RepoNewLesson struct {
	Db *sql.DB
}



func (l *RepoNewLesson) CreateLesson(lesson modul.Lessons) error {

    _, err := l.Db.Exec("insert into lessons (lesson_id, course_id, title, content, created_at) values ($1, $2, $3, $4, $5)",
        lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content, time.Now())
    if err!= nil {
        return err
    }

    return nil
}



func (l *RepoNewLesson) GetAllLesson(filter modul.LessonFilter) (*[]modul.Lessons, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)

	query := `select lesson_id, course_id, title, content, created_at, updated_at from lessons  `

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

	query, arr = ReplaceQueryParamsLesson(query, params)
	rows, err := l.Db.Query(query, arr...)
	if err != nil{
		return nil, err
	}
	lessons := []modul.Lessons{}
	var lesson modul.Lessons
	for rows.Next() {
		err = rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	fmt.Println(lessons)
	return &lessons, nil
}

func ReplaceQueryParamsLesson(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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


func (l *RepoNewLesson) GetLessonById(id string) (*modul.Lessons, error) {
	
	row := l.Db.QueryRow("select lesson_id, course_id, title, content, created_at, updated_at from lessons where lesson_id = $1", id)
	

	lesson := modul.Lessons{}

	err := row.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
	if err!= nil {
        return nil, err
    }
	
	return &lesson, nil
}

func (l *RepoNewLesson) UpdateLesson(lesson modul.Lessons, id string) error {
	
	_, err := l.Db.Exec("Update lessons set title = $1, content = $2, updated_at = $3 where lesson_id = $4",
		lesson.Title, lesson.Content ,time.Now(),id)
	if err != nil {
		return err
	}

	return nil
}

func (l *RepoNewLesson)DeleteLesson(id string) error {
	
	_, err := l.Db.Exec("update lessons set deleted_at = date_part('epoch', current_timestamp)::INT where lesson_id = $1 and deleted_at = 0", id)
	if err != nil {
		return err
	}

	return nil
}

