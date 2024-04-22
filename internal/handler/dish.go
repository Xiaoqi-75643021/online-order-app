package handler

import (
	"net/http"
	"online-ordering-app/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDishes(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")

	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)

	dishes, err := service.ListDishes(pageNum, pageSizeNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取菜品失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取成功", gin.H{"dishes": dishes})
}

func SearchDishes(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		Respond(c, http.StatusBadRequest, 1, "搜索关键词不能为空", nil)
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")

	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)

	dishes, err := service.SearchDishes(keyword, pageNum, pageSizeNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "搜索失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "搜索成功", gin.H{"dishes": dishes})
}

func GetDishesByCategory(c *gin.Context) {
	categoryId := c.Query("categoryId")
	if categoryId == "" {
		Respond(c, http.StatusBadRequest, 1, "分类ID不能为空", nil)
		return
	}
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")

	categoryIdNum, _ := strconv.Atoi(categoryId)
	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)


	dishes, err := service.GetDishesByCategory(uint(categoryIdNum), pageNum, pageSizeNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取分类菜品失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取成功", gin.H{"dishes": dishes})
}

func GetPopularDishes(c *gin.Context) {
	dishes, err := service.GetPopularDishes()
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取热门菜品失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取成功", gin.H{"dishes": dishes})
}

func UpdateDish(c *gin.Context) {
	// Dish字段：Name string	Description string	Price float64	Category string	Ispopular bool
	var dishUpdate map[string]any
	if err := c.BindJSON(&dishUpdate); err != nil {
		Respond(c, http.StatusBadRequest, 1, "参数错误", gin.H{"error": err.Error()})
		return
	}
	dishID := c.Param("id")
	dishIDNum, _ := strconv.Atoi(dishID)
	err := service.UpdateDish(uint(dishIDNum), dishUpdate)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "菜品更新失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "更新成功", nil)
}