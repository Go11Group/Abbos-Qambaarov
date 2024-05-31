package postg

import (
	model "mymod/modul"


	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

func StudentCreate(db *gorm.DB) (*model.Student,error) {
	var user model.Student
	user.Name="Abbos"
	user.Age = 23
	user.Gender = "m"
	user.Course= "Grogramming"
	res := db.Create(&user)
	if res.Error != nil {
		return nil,res.Error
	}
	return &user,nil
	
}

func GetAllStudent(db *gorm.DB ) []model.Student {
	var user []model.Student
    db.Table("students").Select("*").
        Scan(&user)

	return user
    
	
}

func GetStudentByID(db *gorm.DB, id int) (*model.Student, error) {
	var user model.Student
	res := db.First(&user,id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user,nil
}

func StudentUpdate(db *gorm.DB, id int) (*[]model.Student,error) {
	var users []model.Student
	err := db.Model(users).Where("id = ?",id).Update("Course","Dizayn")
	if err.Error != nil {
		return nil,err.Error
	}
	
	return &users,nil
}

func StudentDelete(db *gorm.DB, id int) (*[]model.Student,error) {
	var user []model.Student
	err := db.Where("id = ?", id).Delete(&user)
	if err.Error != nil {
		return nil,err.Error
	}
	return &user,nil
}

