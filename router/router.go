package router

import (
	"TestTask/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.PersonHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		people := api.Group("/people")
		{
			people.POST("/", h.CreatePerson)
			people.GET("/", h.GetPeople)
			people.GET("/:id", h.GetPersonByID)
			people.PUT("/:id", h.UpdatePerson)
			people.DELETE("/:id", h.DeletePerson)
		}
	}

	return r
}
