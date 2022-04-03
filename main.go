package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type unit struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

var units = []unit{
	{ID: "1", Label: "used bike"},
}

func getUnits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, units)
}

func postUnits(c *gin.Context) {
	var newUnit unit

	if err := c.BindJSON(&newUnit); err != nil {
		return
	}

	units = append(units, newUnit)
	c.IndentedJSON(http.StatusCreated, newUnit)
}

func getUnitByID(c *gin.Context) {
	id := c.Param("id")

	for _, u := range units {
		log.Println(u)
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "unit not found"})
}

func main() {
	router := gin.Default()
	router.GET("/units", getUnits)
	router.POST("/units", postUnits)
	router.GET("/units/:id", getUnitByID)

	router.Run("localhost:8080")
}
