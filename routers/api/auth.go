package api

import (
	"net/http"

	"github.com/Crshi/Blog/models"
	"github.com/Crshi/Blog/pkg/e"
	"github.com/Crshi/Blog/pkg/logging"
	"github.com/Crshi/Blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// Param
// name="参数ming"
// value="参数说明"
// dataType="数据类型"
// paramType="query" 表示参数放在哪里
//     · header 请求参数的获取：@RequestHeader
//     · query   请求参数的获取：@RequestParam
//     · path（用于restful接口） 请求参数的获取：@PathVariable
//     · body（不常用）
//     · form（不常用）
// defaultValue="参数的默认值"
// required="true" 表示参数是否必须传

// @Summary Login
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
