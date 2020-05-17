package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"money/models"
	"money/pkg/e"
	"money/pkg/utils"
	"money/pkg/validator"
	"net/http"
	"strconv"
)

//获取所有分类
func GetCategory(c *gin.Context) {
	var category models.Category
	data := models.DB.Find(&category)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": data.Value,
	})
}

//新增分类
func AddCategory(c *gin.Context) {
	var category models.Category
	validate := validator.Validate

	_ = c.ShouldBindJSON(&category)
	err := validate.Struct(category)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		data := models.DB.Create(&category)
		fmt.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//修改分类
func UpdateCategory(c *gin.Context) {
	var category models.Category
	validate := validator.Validate

	_ = c.ShouldBindJSON(&category)
	err := validate.Struct(category)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		data := models.DB.Model(&category).Update(category)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//删除文章标签
func DeleteCategory(c *gin.Context) {
	var category models.Category
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		//res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": "传入的id错误",
		})
	} else {
		data := models.DB.Model(&category).Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}
