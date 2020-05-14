package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"money/models"
	"money/pkg/e"
	"money/pkg/validator"
	"net/http"
)

//获取所有分类
func GetCategory(c *gin.Context) {
}

//新增分类
func AddCategory(c *gin.Context) {
	validate := validator.Validate

	var category models.Category
	_ = c.ShouldBindJSON(&category)
	err := validate.Struct(category)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": err,
		})
	} else {
		data := models.DB.Create(&category)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//修改分类
func UpdateCategory(c *gin.Context) {
	validate := validator.Validate

	var category models.Category
	_ = c.ShouldBindJSON(&category)
	err := validate.Struct(category)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": err,
		})
	} else {
		//category.Name = "cc"
		data := models.DB.Model(&category).Update(category)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//删除文章标签
func DeleteCategory(c *gin.Context) {
}
