package handler

import (
	"fmt"
	models "mymod/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTerminal(ctx *gin.Context) {
	stn := models.Terminals{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Terminal.CreateTerminal(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Created Terminal")
}

func (h *handler) GetTerminalById(ctx *gin.Context) {

	id := ctx.Param("id")
	stn, err := h.Terminal.GetTerminalById(id)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, stn)
}

func (h *handler) GetTerminals(ctx *gin.Context) {

    stns, err := h.Terminal.GetTerminals()
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, stns)
}

func (h *handler) UpdateTerminal(ctc *gin.Context) {

    id := ctc.Param("id")
	stn := models.Terminals{}

    err := ctc.ShouldBindJSON(&stn)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Terminal.UpdateTerminal(stn, id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Updated Terminal")
}

func (h *handler) DeleteTerminal(ctc *gin.Context) {
	id := ctc.Param("id")
	err := h.Terminal.DeleteTerminal(id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Deleted Terminal")
}