package routers

import (
	"github.com/gin-gonic/gin"
	"money/config"
	v1 "money/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(config.GetConfig().RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.POST("/record", v1.AddRecord)
	}

	apiCategory := r.Group("/api/v1/category")
	{
		apiCategory.POST("/add", v1.AddCategory)
		apiCategory.PUT("/update/:id", v1.UpdateCategory)
		apiCategory.GET("/all", v1.GetCategory)
		apiCategory.DELETE("/delete/:id", v1.DeleteCategory)
	}

	return r
}
