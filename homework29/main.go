package main

import (
	"fmt"
	postg "mymod/package"
	

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbURL := "postgres://postgres:root@localhost:5432/education?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userCreate,err := postg.StudentCreate(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(userCreate)

	// allUser := postg.GetAllStudent(db)
	// for _, v := range allUser {

	// 	fmt.Println(v)
	// }

	// IdUser,err := postg.GetStudentByID(db,1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(IdUser)

	// userUpdate, err := postg.StudentUpdate(db, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userUpdate)

	// userDelete, err := postg.StudentDelete(db, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userDelete)


	// courseCreate,err := postg.CourseCreate(db)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(courseCreate)

	// allcourse := postg.GetAllCourse(db)
	// for _, v := range allcourse {

	// 	fmt.Println(v)
	// }

	// Idcourse,err := postg.GetCourseByID(db,1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(Idcourse)

	// courseUpdate, err := postg.CourseUpdate(db, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(courseUpdate)

	// courseDelete, err := postg.CourseDelete(db, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(courseDelete)

}
