package modul

import (
	"database/sql"
	"time"
)

type Courses struct {
	CourseId    string       `json:"course"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   int64        `json:"deleted_at"`
}
type CourseFilter struct {
	Title  string `json:"title"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type Course2 struct {
	CourseId    string       `json:"course"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
}

type CourseLessons struct {
	CourseId string
	Lessons  []Lesson2
}

type CourseUsers struct {
	CourseId string
	Users    []User2
}

