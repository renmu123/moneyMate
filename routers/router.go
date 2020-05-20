package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"money/config"
	v1 "money/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(cors.Default())

	gin.SetMode(config.GetConfig().RunMode)
	// 标签
	apiv1 := r.Group("/api/v1/tag")
	{
		//获取标签列表
		apiv1.GET("/all", v1.GetTags)
		//新建标签
		apiv1.POST("/add", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/update/:id", v1.UpdateTag)
		//删除指定标签
		apiv1.DELETE("/delete/:id", v1.DeleteTag)
	}
	// 分类
	apiCategory := r.Group("/api/v1/category")
	{
		apiCategory.POST("/add", v1.AddCategory)
		apiCategory.PUT("/update/:id", v1.UpdateCategory)
		apiCategory.GET("/all", v1.GetCategory)
		apiCategory.DELETE("/delete/:id", v1.DeleteCategory)
	}
	// 记录
	apiRecord := r.Group("/api/v1/record")
	{
		apiRecord.POST("/add", v1.AddRecord)
		apiRecord.PUT("/delete/:id", v1.UpdateRecord)
		apiRecord.GET("/all", v1.GetRecord)
		apiRecord.DELETE("/:id", v1.DeleteRecord)
	}
	// 账户
	apiAccount := r.Group("/api/v1/acount")
	{
		apiAccount.POST("/add", v1.AddAccount)
		apiAccount.PUT("/delete/:id", v1.UpdateAccount)
		apiAccount.GET("/all", v1.GetAccount)
		apiAccount.DELETE("/:id", v1.DeleteAccount)
	}
	return r
}
