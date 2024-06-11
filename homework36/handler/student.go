package handler

import (
	
	"fmt"
	modul "mymod/module"
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
)

func (s *Handler3) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student := modul.Students{}

	students, err := s.PostStudent.GetAllStudent(student)
	if err != nil {
		panic(err)
	}
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range *students {
		if int(num) == v.Id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{
		"Error": "Not Found Student",
	})

}

func (s *Handler3) GetStudent(c *gin.Context) {
	student := modul.Students{}

	students, err := s.PostStudent.GetAllStudent(student)
	if err != nil {
		panic(err)
	}
	
	for _, v := range *students {
		c.JSON(http.StatusOK,v)
	}

}

func (s *Handler3) CreateStudent(c *gin.Context) {
	student := modul.Students{}

	err := c.BindJSON(&student)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error creating new product")
		return
	}

	s.PostStudent.CreateStudent(student)

	fmt.Println(student)
	c.JSON(http.StatusAccepted,"saved successfully")
}

func (s *Handler3) UpdateStudent(c *gin.Context) {

	id := c.Param("id")
	newStudent := modul.Students{}

	err := c.BindJSON(&newStudent)
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest,"Error entered to read information")
		return
	}
	Student := modul.Students{}

	err = s.PostStudent.UpdateStudent(Student, id)
	c.JSON(http.StatusAccepted,"Saved successfully")

	
}

func (s *Handler3) DeleteStudent(c *gin.Context) {
	
	id := c.Param("id")
	student := modul.Students{}
	err := s.PostStudent.DeleteStudent(student, id)
	if err!= nil {
        panic(err)
    }
	c.JSON(http.StatusAccepted,"DELETE FROM DATABASE")

}
