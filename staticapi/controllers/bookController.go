package controllers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Book{}
	seq   = 1
)

func CreateBooks(c echo.Context) error {
	u := &models.Book{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	books[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBook(c echo.Context) error {
	u := new(models.Book)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Author = u.Author
	books[id].Title = u.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBooks(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.JSON(http.StatusOK, "Success Delete Book")
}

func GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
