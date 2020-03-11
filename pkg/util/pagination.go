package util

import (
	"github.com/Crshi/Blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取SkipCount
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
