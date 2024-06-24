package handler

import (
	"fmt"
	"mymod/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	user := models.CreateUser{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.User.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Created User")
}

func (h *handler) GetUserById(ctx *gin.Context) {

	id := ctx.Param("id")
	user, err := h.User.GetUserById(id)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (h *handler) GetUsers(ctx *gin.Context) {

    users, err := h.User.GetUser()
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, users)
}

func (h *handler) UpdateUser(ctc *gin.Context) {

    id := ctc.Param("id")
	user := models.CreateUser{}

    err := ctc.ShouldBindJSON(&user)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.User.UpdateUser(user, id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Updated User")
}

func (h *handler) DeleteUser(ctc *gin.Context) {
	id := ctc.Param("id")
	err := h.User.DeleteUser(id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Deleted User")
}
