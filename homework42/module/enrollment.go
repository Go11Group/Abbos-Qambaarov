package modul

import (
	"database/sql"
	"time"
)

type Enrollments struct {
	EnrollmentId   string       `json:"enrollment_id"`
	UserId         string       `json:"user_id"`
	CourseId       string       `json:"course_id"`
	EnrollmentDate time.Time    `json:"enrollment_date"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
	DeletedAt      int64        `json:"deleted_at"`
}

type EnrollmentFilter struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
