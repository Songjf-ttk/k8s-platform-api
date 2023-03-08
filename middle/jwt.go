package middle

import (
	"k8s-plantform/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//对登录接口放行
		if len(c.Request.URL.String()) >= 10 && c.Request.URL.String()[0:10] == "/api/login" {
			c.Next()
		} else {
			//处理验证逻辑
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  "请求未携带token，无权限访问",
					"data": nil,
				})
				c.Abort()
				return
			}
			//解析token内容
			claims, err := utils.JWTToken.ParseToken(token)
			if err != nil {
				//token过期错误
				if err.Error() == "TokenExpired" {
					c.JSON(http.StatusBadRequest, gin.H{
						"msg":  "授权已过期",
						"data": nil,
					})
					c.Abort()
					return
				}
				//其他解析错误
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  err.Error(),
					"data": nil,
				})
				c.Abort()
				return
			}
			c.Set("claims", claims)
			c.Next()
		}
	}
}
