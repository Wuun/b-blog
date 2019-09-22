package service

import (
	"bblog/conf"
	"bblog/serializer"

	"github.com/pborman/uuid"
)

//UploadArticleSrv is the set of mathods of upload article.
type UploadArticleSrv struct {
	Title       string `form:"title" json:"title"`
	Tag         string `form:"tag" json:"tag"`
	Description string `form:"description" json:"description"`
	Content     string `form:"content" json:"content"`
}

//Upload use to upload article.
func (service *UploadArticleSrv) Upload() *serializer.Response {
	if service.Description == "" || service.Title == "" || service.Content == "" {
		return &serializer.Response{
			Status: 40003,
			Msg:    "upload failed,please input completed.",
		}
	}
	ud := uuid.New()
	articleModel := &serializer.ArticleModel{
		UUID:        ud,
		Tag:         service.Tag,
		Description: service.Description,
		Title:       service.Title,
	}
	articleContentModel := &serializer.ArticleContent{
		ArticleModelID: ud,
		Content:        service.Content,
	}

	if err := conf.MYSQL_CONNECT.Create(articleModel).Error; err != nil {
		return &serializer.Response{
			Status: 40002,
			Msg:    "can't not insert this article into server,error-1.",
			Error:  err,
		}
	}

	if err := conf.MYSQL_CONNECT.Create(articleContentModel).Error; err != nil {
		conf.MYSQL_CONNECT.Delete(articleModel).Where("uuid = ?", ud)
		return &serializer.Response{
			Status: 40002,
			Msg:    "can't not insert this article into server.error-2",
			Error:  err,
		}
	}

	return &serializer.Response{
		Status: 0,
		Msg:    "success",
	}
}
