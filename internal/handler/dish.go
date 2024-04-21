package handler

import (
	"net/http"
	"online-ordering-app/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchDishes(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		Respond(c, http.StatusBadRequest, 1, "搜索关键词不能为空", nil)
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

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
	pageSize := c.DefaultQuery("pageSize", "10")

	categoryIdNum, _ := strconv.Atoi(categoryId)
	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)


	dishes, err := service.GetDishesByCategory(uint(categoryIdNum), pageNum, pageSizeNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取菜品失败", gin.H{"error": err.Error()})
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
