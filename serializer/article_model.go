package serializer

import "github.com/jinzhu/gorm"

//ArticleModel is the gorm model of article
type ArticleModel struct {
	gorm.Model
	UUID        string `gorm:"UNIQUE;type:varchar(255)"`
	Tag         string `json:"tag" gorm:"type:VARCHAR(100)"`
	Title       string `json:"title" gorm:"type:VARCHAR(100)"`
	Description string `gorm:"type:VARCHAR(100)" json:"description"`
}

//ArticleContent is the gorm model of article`s content
type ArticleContent struct {
	ArticleModel   ArticleModel `json:"art_model" gorm:"foreignkey:ArticleModelID"`
	ArticleModelID string       `gorm:"PRIMARY_KEY"`
	Content        string       `gorm:"type:text;" json:"content"`
}

//Comment is the comment of a specific article.
type Comment struct {
	ArticleContent ArticleContent `gorm:"foreignkey:ArticleID"`
	ArticleID      string
	Text           string `gorm:"VARCHAR(255)"`
	Author         string `gorm:"type:VARCHAR(20)"`
	IP             string `gorm:"type:VARCHAR(20)"`
}
