package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"money/models"
)

func AddRecord(c *gin.Context) {
	var record models.Record
	err := c.BindJSON(&record)
	fmt.Println(record.Name)
	fmt.Println(record.Money)
	if err != nil {
		log.Fatal(err)
		//record := models.Record{Name: record.Name, Type: record.Type, Money: record.Money, Desc: record.Desc}
	}
	models.DB.Create(&record)
	c.JSON(200, gin.H{
		"he": "jhello",
	})

}
