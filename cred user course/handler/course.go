package handler

import (
	"fmt"
	modul "mymod/module"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handlerdan get course by id API ni ishlatish
func (co *Handler1) GetCourseByID(c *gin.Context) {
	//urldan idni olish
	id := c.Param("id")

	// postgresdan get course by id funksiyasini chaqirish
	course, err := co.PostCourse.GetCourseById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Invalid limit parameter")
	}

	if id == course.CourseId {
		
		c.JSON(http.StatusOK,course)
		return
	}
	
	c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found Course"})
		

}
// handlerdan get All course API ni ishlatish
func (co *Handler1) GetAllCourse(c *gin.Context) {

	filter := modul.CourseFilter{}
	filter.Title = c.Param("title")

	lim := c.Param("limit")
	if lim != "" {
		limit, err := strconv.Atoi(lim)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Invalid limit parameter")
			return
		}
		filter.Limit = limit
	}

	offset := c.Param("offset")
	if offset != "" {
		offs, err := strconv.Atoi(offset)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Invalid offset parameter")
			return
		}
		filter.Offset = offs
	}
	
	courses, err := co.PostCourse.GetAllCourse(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK,courses)
	
}

// handlerdan Create course API ni ishlatish
func (co *Handler1) CreateCourse(c *gin.Context) {
	
	course := modul.Courses{}
	err := c.BindJSON(&course)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error Create Course")
		return
	}
	fmt.Println(course)
	err = co.PostCourse.CreateCourse(course)
	if err != nil {
		fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "Error creating server")
        return
	}
	fmt.Println(course)
	c.JSON(http.StatusAccepted,"saved successfully")
}

// handlerdan update course ni ishlatish
func (co *Handler1) UpdateCourse(c *gin.Context) {

	id :=c.Param("id")
	NewCourse := modul.Courses{}

	err := c.BindJSON(&NewCourse)
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest,"Error update user")
		return
	}

	err = co.PostCourse.UpdateCourse(NewCourse,id)
	if err != nil {
		fmt.Println("Error update",err)
        c.JSON(http.StatusInternalServerError,"Error Update in database")
        return
	}
	c.JSON(http.StatusAccepted,"Saved successfully")
}

// handlerdan delete course API ni ishlatish
func (co *Handler1) DeleteCourse(c *gin.Context) {
	//urldan idni olish
	id := c.Param("id")

	err := co.PostCourse.DeleteCourse(id)
	if err!= nil {
        c.JSON(http.StatusInternalServerError,"Error Delete")
    }
	c.JSON(http.StatusAccepted,"DELETE FROM DATABASE")

}
