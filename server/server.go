package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miketonks/swag"
	"github.com/miketonks/swag-validator"
	"github.com/miketonks/swag/endpoint"
	"github.com/miketonks/swag/swagger"
)

func SetupAPI() *swagger.API {
	addProvidable := endpoint.New("post", "/providable", "add a new providable to the db",
		endpoint.Handler(AddProvidable),
		endpoint.Description(""),
		endpoint.Body(Providable{}, "providable to be added to the db", true),
		endpoint.Response(http.StatusOK, Providable{}, "successfully added providable"),
		endpoint.Tags("providable", "config"),
	)

	api := swag.New(swag.Endpoints(addProvidable))
	return api
}

func SetupRouter(api *swagger.API) *gin.Engine {
	router := gin.New()
	enableCors := true
	router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))

	router.Use(swagvalidator.SwaggerValidator(api))

	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		colonPath := swag.ColonPath(path)

		router.Handle(endpoint.Method, colonPath, h)
	})
	return router
}

type Providable struct {
	ID         int64        `json:"id"`
	UUID       swagger.UUID `json:"uuid"`
	Name       string       `json:"name" binding:"required"`
	CreateTime time.Time    `json:"create_time" binding:"required"`
}

func AddProvidable(c *gin.Context) {
	var providable Providable
	if err := c.ShouldBindJSON(&providable); err == nil {
		c.JSON(http.StatusOK, providable)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
