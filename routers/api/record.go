package v1

import (
	"github.com/gin-gonic/gin"
	"money/models"
	"money/pkg/e"
	"money/pkg/utils"
	"money/pkg/validator"
	"net/http"
	"strconv"
)

func AddRecord(c *gin.Context) {
	validate := validator.Validate

	var record models.Record
	_ = c.ShouldBindJSON(&record)
	err := validate.Struct(record)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		models.DB.Create(&record)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": "",
		})
	}
}

func GetRecord(c *gin.Context) {
	var record models.Record
	data := models.DB.Find(&record)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": data.Value,
	})
}

func DeleteRecord(c *gin.Context) {
	var record models.Record
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  0,
			"error": "传入的id错误",
		})
	} else {
		data := models.DB.Model(&record).Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}

}

func UpdateRecord(c *gin.Context) {
	var record models.Record
	validate := validator.Validate

	_ = c.ShouldBindJSON(&record)
	err := validate.Struct(record)
	if err != nil {
		res := utils.HandleCommonError(err)
		c.JSON(400, gin.H{
			"code":  0,
			"error": res,
		})
	} else {
		data := models.DB.Model(&record).Update(record)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": data.Value,
		})
	}
}
