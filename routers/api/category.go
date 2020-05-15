package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
	"money/models"
	"money/pkg/e"
	"money/pkg/validator"
	"net/http"
	"strings"
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
		res := CommonError(err)
		c.JSON(400, gin.H{
			"code":  10,
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

func CommonError(err error) map[string]string {
	errs := err.(v.ValidationErrors)
	res := make(map[string]string)
	for _, er := range errs {
		f := strings.ToLower(er.Field())
		fmt.Println(f)
		res[f] = er.Tag()
	}
	return res
}

//删除文章标签
func DeleteCategory(c *gin.Context) {
}
