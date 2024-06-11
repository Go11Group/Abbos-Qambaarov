package handler

import (
	"fmt"
	modul "mymod/module"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (u *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user := modul.Users{}

	users, err := u.PostUser.GetAllUser(user)
	if err != nil {
		panic(err)
	}

	for _,v := range *users {
		if id == v.Id {
			
			c.JSON(http.StatusOK,v)
			return

		}
	}
	c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found User"})
		

}

func (u *Handler) GetUser(c *gin.Context) {
	
	user := modul.Users{}
	users, err := u.PostUser.GetAllUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
		c.JSON(http.StatusOK,users)
	
}

func (u *Handler) CreateUser(c *gin.Context) {
	user := modul.Users{}

	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error entered to read information")
		return
	}

	err = u.PostUser.CreateUser(user)
	if err != nil {
		fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "Error creating new product")
        return
	}
	fmt.Println(user)
	c.JSON(http.StatusAccepted,"saved successfully")
}

func (u *Handler) UpdateUser(c *gin.Context) {

	id :=c.Param("id")
	newuser := modul.Users{}

	err := c.BindJSON(&newuser)
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest,"Error update user")
		return
	}
	user := modul.Users{}

	err = u.PostUser.UpdateUser(user,id)
	if err != nil {
		fmt.Println("Error update",err)
        c.JSON(http.StatusInternalServerError,"Error Update in database")
        return
	}
	c.JSON(http.StatusAccepted,"Saved successfully")
}

func (u *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	user := modul.Users{}
	err := u.PostUser.DeleteUser(user,id)
	if err!= nil {
        panic(err)
    }
	c.JSON(http.StatusAccepted,"DELETE FROM DATABASE")

}
