package postg

import (
	model "mymod/modul"


	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

func CourseCreate(db *gorm.DB) (*model.Course,error) {
	var user model.Course
	user.Name="Dizayn"
	user.Field= "Grogramming"
	res := db.Create(&user)
	if res.Error != nil {
		return nil,res.Error
	}
	return &user,nil
	
}

func GetAllCourse(db *gorm.DB ) []model.Course {
	var user []model.Course
    db.Table("Course").Select("*").
        Scan(&user)

	return user
    
	
}

func GetCourseByID(db *gorm.DB, id int) (*model.Course, error) {
	var user model.Course
	res := db.First(&user,id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user,nil
}

func CourseUpdate(db *gorm.DB, id int) (*[]model.Course,error) {
	var users []model.Course
	err := db.Model(users).Where("id = ?",id).Update("Course","Dizayn")
	if err.Error != nil {
		return nil,err.Error
	}
	
	return &users,nil
}

func CourseDelete(db *gorm.DB, id int) (*[]model.Course,error) {
	var user []model.Course
	err := db.Where("id = ?", id).Delete(&user)
	if err.Error != nil {
		return nil,err.Error
	}
	return &user,nil
}

