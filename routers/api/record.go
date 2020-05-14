package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"money/models"
	"money/pkg/e"
	"money/pkg/validator"
	"net/http"
)

func AddRecord(c *gin.Context) {
	validate := validator.Validate

	var record models.Record
	_ = c.ShouldBindJSON(&record)
	err := validate.Struct(record)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"errow": err,
		})
	} else {
		models.DB.Create(&record)
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": "",
		})
	}

}

func DeleteRecord(c *gin.Context) {

}

func UpdateRecord(c *gin.Context) {

}
