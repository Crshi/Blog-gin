package routers

import (
	"github.com/Crshi/Blog/pkg/setting"
	v1 "github.com/Crshi/Blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	//初始化中间件
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
