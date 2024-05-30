package main

import (
	// "fmt"
	// "mymod/model"
	"mymod/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	
	// err = postgres.StudentCreate(db)
	// if err != nil {
	// 	panic(err)
	// }
	
	// err = postgres.CourseCreate(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = postgres.StudentInsert(db)
	// if err != nil {
	// 	panic(err)
	// }
	
	// err = postgres.CourseInsert(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = postgres.StudentUpdate(db)
	// if err != nil {
	// 	panic(err)
	// }
	
	// err = postgres.CourseUpdate(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = postgres.StudentDelete(db)
	// if err != nil {
	// 	panic(err)
	// }
	
	// err = postgres.CourseDelete(db)
	// if err != nil {
	// 	panic(err)
	// }

	

	// st := postgres.NewStudentRepo(db)

	// users, err := st.GetAllStudents()
	// if err != nil {
	// 	panic(err)
	// }

	// user, _ := st.GetByStusentID(users[2].ID)

	// fmt.Println(users)

	// fmt.Println(user)

	// cr := postgres.NewCourseRepo(db)
	// courses, err := st.GetAllCourses()
	// if err != nil {
	// 	panic(err)
	// }

	// course, _ := cr.GetByCoursesID(users[2].ID)

	// fmt.Println(courses)

	// fmt.Println(course)


}
