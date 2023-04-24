package controller

import (
	"net/http"
	"strconv"

	"deploy/database"
	"deploy/model"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []model.Blog

	if err := database.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	blog := model.Blog{}
	blogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.Preload("User").First(&blog, blogId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blog":    blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	c.Bind(&blog)

	if err := database.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	blog := model.Blog{}
	blogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.Delete(&blog, blogId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "blog deleted successfully",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	blog := model.Blog{}
	blogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&blog, blogId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	c.Bind(&blog)
	var temp model.Blog = model.Blog{
		Judul:  blog.Judul,
		Konten: blog.Konten,
	}

	if err1 := database.DB.Save(&temp).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "blog updated successfully",
	})
}
