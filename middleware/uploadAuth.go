package middleware

import (
	"bblog/conf"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UploadAuth is middleware use to judge use is admin or not.
func UploadAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginPage := "http://" + conf.G_CONF.Domain + "/api/v1/login"
		session := sessions.Default(c)
		pas := session.Get("password")
		fmt.Println(conf.G_CONF.WebsitePassWord)
		if pas != nil {
			if pas.(string) == conf.G_CONF.WebsitePassWord {
				c.Next()
				return
			}
			fmt.Println(pas.(string))
			c.Redirect(307, loginPage)
			c.Abort()
		}
		c.Redirect(307, loginPage)
		c.Abort()
		return
	}
}
