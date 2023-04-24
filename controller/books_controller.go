package controller

import (
	"net/http"
	"strconv"

	"deploy/database"
	"deploy/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	book := model.Book{}
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&book, bookId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := database.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	book := model.Book{}
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.Delete(&book, bookId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "book deleted successfully",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	book := model.Book{}
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&book, bookId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	c.Bind(&book)
	if err1 := database.DB.Save(&book).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "book updated successfully",
	})
}
