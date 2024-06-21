package handler

import (
	"fmt"
	modul "mymod/module"
	packages "mymod/storage/postgres"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handlerdan get user by id API ni ishlatish
func (u *Handler) GetUserByID(c *gin.Context) {

	// urldan idni olish
	id := c.Param("id")

	// posgresdan get by idni chaqirish va errorga tekshirish
	users, err := u.PostUser.GetUserById(id)
	if err != nil {
		panic(err)
	}

	if id == users.UserId {
		//postmenda elon qilish
		c.JSON(http.StatusOK, users)
		return
	}
	// xatolikni elon qilish
	c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found User"})

}

// handlerdan get All users API ni ishlatish
func (u *Handler) GetUsers(c *gin.Context) {
	filter := modul.Filter{}
	// urldan idni olish
	filter.Name = c.Param("name")

	// urldan ageni olish va Atoi orqali intga o'tkazish
	num := c.Param("age")
	if num != "" {
		nums, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Invalid age parameter")
			return
		}
		filter.Age = nums
	}
	// urldan emailni olish
	filter.Email = c.Param("email")

	// urldan limitni olish va Atoi orqali intga o'tkazish
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

	// urldan offsetni olish va Atoi orqali intga o'tkazish
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

	// posgresdan get All userni chaqirish va errorga tekshirish
	users, err := u.PostUser.GetAllUser(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// postmenda elon qilish
	c.JSON(http.StatusOK, users)
}

// handlerdan create users API ni ishlatish
func (u *Handler) CreateUser(c *gin.Context) {
	user := modul.Users{}

	// bodydan structni olish va xatolikni tekshirish
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		// xatolikni elon qilish
		c.JSON(http.StatusBadRequest, "Error Creating user")
		return
	}

	// postgresdan create userni chaqirish
	err = u.PostUser.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		// xatolikni elon qilish
		c.JSON(http.StatusInternalServerError, "Error server")
		return
	}
	fmt.Println(user)
	// postmenda elon qilish
	c.JSON(http.StatusOK, "saved successfully")
}

// handlerdan update users API ni ishlatish
func (u *Handler) UpdateUser(c *gin.Context) {

	//urldan idni olish
	id := c.Param("id")
	newUser := modul.Users{}

	//bodydan structni olish
	err := c.BindJSON(&newUser)
	// errorni tekshirish
	if err != nil {
		fmt.Println("Error entered to read information")
		c.JSON(http.StatusBadRequest, "Error update user")
		return
	}

	//postgresdan update userni chaqirish
	err = u.PostUser.UpdateUser(newUser, id)
	// errorni tekshirish
	if err != nil {
		fmt.Println("Error update", err)
		//errorni elon qilish
		c.JSON(http.StatusInternalServerError, "Error Update in database")
		return
	}
	// postmenda elon qilish
	c.JSON(http.StatusAccepted, "Saved successfully")
}

// handlerdan delete user API ni ishlatish
func (u *Handler) DeleteUser(c *gin.Context) {
	// urldan emailni olish
	id := c.Param("id")

	err := u.PostUser.DeleteUser(id)
	// errorni tekshirish
	if err != nil {
		// xatolikni elon qilish
		c.JSON(http.StatusInternalServerError, "Error Delete")
	}
	// postmenda elon qilish
	c.JSON(http.StatusAccepted, "DELETE FROM DATABASE")

}

// handlerdan userni courslarini API sini ishlatish
func (u *Handler) GetUserCourses(c *gin.Context) {
	// urldan emailni olish
	id := c.Param("user_id")
	user, err := u.PostUser.GetUserCourses(id)
	// errorni tekshirish
	if err != nil {
		// xatolikni elon qilish
		err = fmt.Errorf("error on Get user courses: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	}
	// postmenda elon qilish
	c.JSON(http.StatusAccepted, *user)

}

// handlerdan gsearch users API ni ishlatish

func (u *Handler) SearchUsers(c *gin.Context) {
	filter := packages.SearchFilter{}
	// errorni tekshirish
	if err := c.ShouldBindQuery(&filter); err != nil {
		// xatolikni elon qilish
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//postgresdan search userni chaqirish
	users, err := u.PostUser.SearchUsers(filter)
	// errorni tekshirish
	if err != nil {
		// xatolikni elon qilish

		c.Writer.Write([]byte(err.Error()))
	}
	// postmenda elon qilish
	c.JSON(http.StatusOK, *users)
}
