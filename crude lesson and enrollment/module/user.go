package modul

import (
	"database/sql"
	"time"
)

type Users struct {
	UserId    string       `json:"user_id"`
	Name      string       `json:"name"`
	Age       int          `json:"age"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeleteAt  int64        `json:"delete_at"`
}

type Filter struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type UserCourses struct {
	UserId  string    
	Courses []Course2
}

type User2 struct {
	UserID string
	Name   string 
	Email  string 
}

type SearchUser struct {
	Results []User2
   }