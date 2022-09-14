package routes

import (
	"main/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetSingleUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.EditUserController)
	e.POST("/users", controllers.CreateUserController)
	return e
	// materi
	// https://docs.google.com/presentation/d/16Nk4M_B07koXb2kKRIWC4xjrUtr3FkXLMPWEckFYczg/edit#slide=id.gb869e6a5bc_0_403

	// tugas
	// https://docs.google.com/document/d/1_rL66GAW95VdmcMqzVsVLavD0Di3fRcXpnsFKK5cXYE/edit
}
