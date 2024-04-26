package handler

import (
	"net/http"
	"online-ordering-app/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	type request struct {
		Catetory string `json:"category" binding:"required"`
		ParentID *uint  `json:"parent_id"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}
	err := service.Addategory(req.Catetory, req.ParentID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "分类添加失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "分类添加成功", nil)
}

func RemoveCategory(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}
	err := service.RemoveCategory(req.ID)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "分类删除失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "分类删除成功", nil)
}

func GetAllCategories(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")

	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)

	categories, err := service.ListCategories(pageNum, pageSizeNum)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "获取分类失败", gin.H{"error": err.Error()})
		return
	}

	Respond(c, http.StatusOK, 0, "分类获取成功", gin.H{"categories": categories})
}
