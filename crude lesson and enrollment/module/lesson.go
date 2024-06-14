package modul

import (
	"database/sql"
	"time"
)

type Lessons struct {
	LessonId  string       `json:"lesson_id"`
	CourseId  string       `json:"course_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt int64        `json:"deleted_at"`
}

type LessonFilter struct {
	Title  string `json:"title"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type Lesson2 struct {
	LessonId  string       `json:"lesson_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
}