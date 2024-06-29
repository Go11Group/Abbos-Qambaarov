package handler

import (
	"fmt"
	models "mymod/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCard(ctx *gin.Context) {
	card := models.Cards{}

	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Card.CreateCard(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Created Card")
}

func (h *handler) GetCardById(ctx *gin.Context) {

	id := ctx.Param("id")
	card, err := h.Card.GetCardById(id)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, card)
}

func (h *handler) GetCards(ctx *gin.Context) {

    cards, err := h.Card.GetCard()
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, cards)
}

func (h *handler) UpdateCard(ctc *gin.Context) {

    id := ctc.Param("id")
	card := models.Cards{}

    err := ctc.ShouldBindJSON(&card)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Card.UpdateCard(card, id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Updated Card")
}

func (h *handler) DeleteCard(ctc *gin.Context) {
	id := ctc.Param("id")
	err := h.Card.DeleteCard(id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Deleted Card")
}