package handler

import (
	"fmt"
	modul "mymod/module"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *Handler3) GetEnrollmentByID(c *gin.Context) {

	id := c.Param("id")

	enrol, err := e.PostEnrollment.GetEnrollmentById(id)
	if err != nil {
		panic(err)
	}

	if id == enrol.EnrollmentId {
		
		c.JSON(http.StatusOK,enrol)
		return
	}
	
	c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found User"})
		

}

func (e *Handler3) GetAllEnrollment(c *gin.Context) {

	filter := modul.EnrollmentFilter{}
	
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
	
	
	enrollments, err := e.PostEnrollment.GetAllEnrolment(filter)
	fmt.Println("err :>>>>>>>>>> ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK,enrollments)
	
}

func (e *Handler3) CreateEnrollment(c *gin.Context) {
	
	enrol := modul.Enrollments{}
	err := c.BindJSON(&enrol)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error create enrollment")
		return
	}

	fmt.Println(enrol)
	err = e.PostEnrollment.CreateEnrollment(enrol)
	if err != nil {
		fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "Error creating enrolment")
        return
	}
	fmt.Println(enrol)
	c.JSON(http.StatusAccepted,"saved successfully")
}

func (e *Handler3) UpdateEnrollment(c *gin.Context) {

	id :=c.Param("id")

	err := e.PostEnrollment.UpdateEnrollment(id)
	if err != nil {
		fmt.Println("Error update",err)
        c.JSON(http.StatusInternalServerError,"Error Update in database")
        return
	}
	c.JSON(http.StatusAccepted,"Saved successfully")
}

func (e *Handler3) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")

	err := e.PostEnrollment.DeleteEnrollment(id)
	if err!= nil {
        c.JSON(http.StatusInternalServerError,"Error Delete")
    }
	c.JSON(http.StatusAccepted,"DELETE FROM DATABASE")

}