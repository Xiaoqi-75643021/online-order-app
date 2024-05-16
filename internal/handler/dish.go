package handler

import (
	"net/http"
	"online-ordering-app/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddDish(c *gin.Context) {
	type request struct {
		Name     string  `json:"name" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Catetory string  `json:"category" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}
	err := service.AddDish(req.Name, req.Price, req.Catetory)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "菜品添加失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "菜品添加成功", nil)
}

func RemoveDish(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}
	err := service.DeleteDish(req.ID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "菜品删除失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "菜品删除成功", nil)
}

func QueryDishInfoById(c *gin.Context) {
	dishId := c.Query("id")
	if dishId == "" {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}

	dishIdNum, _ := strconv.Atoi(dishId)

	dish, err := service.GetDishInfo(dishIdNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取菜品详情失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取详情成功", gin.H{"dishInfo": dish})
}


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

	Respond(c, http.StatusOK, 0, "菜品获取成功", gin.H{"dishes": dishes})
}

func UpdateDish(c *gin.Context) {
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
	categoryId := c.Query("category_id")
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

	Respond(c, http.StatusOK, 0, "获取分类菜品成功", gin.H{"dishes": dishes})
}

func GetPopularDishes(c *gin.Context) {
	dishes, err := service.GetPopularDishes()
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取热门菜品失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "获取热门菜品成功", gin.H{"dishes": dishes})
}