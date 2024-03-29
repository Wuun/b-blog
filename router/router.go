package router

import (
	"bblog/api"
	"bblog/conf"
	"bblog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

//NewRouter is thr factory of router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	group := router.Group("/api/v1")
	group.Use(middleware.Session(conf.G_CONF.Secret))
	{
		group.StaticFS("/static",http.Dir("./view"))
		group.GET("/login", api.Login)
		group.POST("/authenticate", api.Authenticate)
		group.GET("/ping", api.Ping)
		group.GET("/home", api.ListArticle)
		group.GET("/article_content/:article_id", api.GetArticleContent)
		group.POST("/comment/:article_id", api.CommentToArticle)

		authen := group.Group("write_article/")
		authen.Use(middleware.UploadAuth())
		{
			authen.GET("write", api.WriteArticle)
			authen.POST("upload_article", api.UploadArticle)
		}
	}

	return router
}
