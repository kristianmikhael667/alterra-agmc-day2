package routes

import (
	"main/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.DELETE("/books/:id", controllers.DeleteBooks)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.POST("/books", controllers.CreateBooks)
	return e
}
