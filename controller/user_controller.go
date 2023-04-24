package controller

import (
	"net/http"
	"strconv"

	"deploy/database"
	"deploy/model"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := database.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	user := model.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	user := model.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.Delete(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := model.User{}
	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&user, UserId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	c.Bind(&user)
	if err1 := database.DB.Save(&user).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}
