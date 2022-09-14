package controllers

import (
	"main/lib/database"
	"main/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All users",
		"data":    users,
	})
}

func GetSingleUserController(c echo.Context) error {
	id := c.Param("id")

	users, err := database.GetSingleUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get users",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Name is required",
		})

	} else if user.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email is required",
		})
	} else if user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password is required",
		})
	} else {
		users, err := database.CreateUser(user.Name, user.Email, user.Password)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create New User",
			"data":    users,
		})
	}

}

func EditUserController(c echo.Context) error {
	user := models.User{}
	id := c.Param("id")

	c.Bind(&user)

	users, err := database.EditUser(id, user.Name, user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Edit users success",
		"data":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	users, err := database.DeleteUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Delete users success",
		"data":    users,
	})
}
