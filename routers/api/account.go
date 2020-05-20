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
func GetAccount(c *gin.Context) {
	var account models.Account
	data := models.DB.Find(&account)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": data.Value,
	})
}

//新增分类
func AddAccount(c *gin.Context) {
	var account models.Account
	validate := validator.Validate

	_ = c.ShouldBindJSON(&account)
	err := validate.Struct(account)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		data := models.DB.Create(&account)
		fmt.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//修改分类
func UpdateAccount(c *gin.Context) {
	var account models.Account
	validate := validator.Validate

	_ = c.ShouldBindJSON(&account)
	err := validate.Struct(account)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		data := models.DB.Model(&account).Update(account)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}

//删除文章标签
func DeleteAccount(c *gin.Context) {
	var account models.Account
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  0,
			"error": "传入的id错误",
		})
	} else {
		data := models.DB.Model(&account).Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}
