package v1

import (
	"github.com/gin-gonic/gin"
	"money/models"
)

func AddRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(400, gin.H{
			"errow": err,
		})
		//record := models.Record{Name: record.Name, Type: record.Type, Money: record.Money, Desc: record.Desc}
	} else {
		models.DB.Create(&record)
		c.JSON(200, gin.H{
			"he": "jhello",
		})
	}

}
