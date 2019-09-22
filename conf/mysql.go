package conf

import (
	"bblog/serializer"
	"fmt"

	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//MYSQL_CONNECT is singlten of mysql connection.
var MYSQL_CONNECT *gorm.DB

func initMySQL() {
	var (
		err error
	)
	sqlAdr := G_CONF.MSUser + ":" + G_CONF.MSPWD + `@tcp(` + G_CONF.MSHost + ":" + G_CONF.MSPort + `)/` + G_CONF.MSDBName + "?charset=utf8&parseTime=true"
	fmt.Println(sqlAdr)
	fmt.Println("here")
	MYSQL_CONNECT, err = gorm.Open("mysql", sqlAdr)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect MySQL successfully.")
	Migrate()
	MYSQL_CONNECT.DB().SetMaxIdleConns(20)
	MYSQL_CONNECT.DB().SetMaxOpenConns(100)
}

//Migrate is auto magirate table of server.
func Migrate() {
	if err := MYSQL_CONNECT.AutoMigrate(&serializer.ArticleModel{}).Error; err != nil {
		panic("error,when try to create table")
	}
	if err := MYSQL_CONNECT.AutoMigrate(&serializer.ArticleContent{}).Error; err != nil {
		fmt.Println("error,when try to migrate databases.")
		panic(err)
	}
	if err := MYSQL_CONNECT.AutoMigrate(&serializer.Comment{}).Error; err != nil {
		fmt.Println("error,when try to migrate databases.")
		panic(err)
	}
}
