package server

import (
	"humdrum/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addCar(ctx *gin.Context) {
	car := new(store.Car)
	if err := ctx.Bind(car); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddCar(car); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Added successfully.",
	})
}

func getCars(ctx *gin.Context) {
	partialCar := new(store.PartialCar)
	if err := ctx.Bind(partialCar); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matchingCars, err := store.GetMatchingCars(partialCar)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, matchingCars)
}

func updateCar(ctx *gin.Context) {
	partialCarUpdate := new(store.PartialCar)

	if err := ctx.Bind(partialCarUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCar, err := store.UpdateCar(partialCarUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedCar)
}

func deleteCar(ctx *gin.Context) {
	// Extract the car ID from the request parameters.
	carIDStr := ctx.Param("carId")
	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = store.DeleteCar(carID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
