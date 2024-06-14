package handler

import (
	"fmt"
	modul "mymod/module"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (l *Handler2) GetLessonByID(c *gin.Context) {

	id := c.Param("id")

	lesson, err := l.PostLesson.GetLessonById(id)
	if err != nil {
		panic(err)
	}

	if id == lesson.LessonId {
		
		c.JSON(http.StatusOK,lesson)
		return
	}
	
	c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found User"})
		

}

func (l *Handler2) GetAllLesson(c *gin.Context) {

	filter := modul.LessonFilter{}
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
	
	lessons, err := l.PostLesson.GetAllLesson(filter)
	if err != nil {
		fmt.Println("error>>>>>>>>>>>>>",err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, lessons)
	
}

func (l *Handler2) CreateLesson(c *gin.Context) {
	
	lesson := modul.Lessons{}
	err := c.ShouldBind(&lesson)
	if err != nil {
		fmt.Println(err)	
		c.JSON(http.StatusBadRequest, "Error create lesson")
		return
	}

	fmt.Println(lesson)
	err = l.PostLesson.CreateLesson(lesson)
	if err != nil {
		fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "Error creating server")
        return
	}
	fmt.Println(lesson)
	c.JSON(http.StatusAccepted,"saved successfully")
}

func (l *Handler2) UpdateLesson(c *gin.Context) {

	id :=c.Param("id")
	NewLesson := modul.Lessons{}

	err := c.BindJSON(&NewLesson)
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest,"Error update user")
		return
	}

	err = l.PostLesson.UpdateLesson(NewLesson,id)
	if err != nil {
		fmt.Println("Error update",err)
        c.JSON(http.StatusInternalServerError,"Error Update in database")
        return
	}
	c.JSON(http.StatusAccepted,"Saved successfully")
}

func (l *Handler2) DeleteLesson(c *gin.Context) {
	id := c.Param("id")

	err := l.PostLesson.DeleteLesson(id)
	if err!= nil {
        c.JSON(http.StatusInternalServerError,"Error Delete")
    }
	c.JSON(http.StatusAccepted,"DELETE FROM DATABASE")

}