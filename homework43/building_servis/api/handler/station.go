package handler

import (
	"fmt"
	"mymod/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	stn := models.CreateStation{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.CreateStation(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "Created Station")
}

func (h *handler) GetStationById(ctx *gin.Context) {

	id := ctx.Param("id")
	stn, err := h.Station.GetStationById(id)
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, stn)
}

func (h *handler) GetStations(ctx *gin.Context) {

    stns, err := h.Station.GetStation()
    if err!= nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("Error from server:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, stns)
}

func (h *handler) UpdateStation(ctc *gin.Context) {

    id := ctc.Param("id")
	stn := models.CreateStation{}

    err := ctc.ShouldBindJSON(&stn)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Station.UpdateStation(stn, id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Updated Station")
}

func (h *handler) DeleteStation(ctc *gin.Context) {
	id := ctc.Param("id")
	err := h.Station.DeleteStation(id)
    if err!= nil {
        ctc.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctc.JSON(http.StatusOK, "Deleted Station")
}
