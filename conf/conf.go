package conf

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

//G_CONF is config will be used.
var G_CONF *Conf

//Conf is conf
type Conf struct {
	MaxConnected    int
	ConnectTimeout  time.Duration
	MSHost          string
	MSPWD           string
	MSPort          string
	MSDBName        string
	MSUser          string
	WebsitePassWord string
	WebAddr         string
	Secret          string
	Domain          string
	StaticAddr      string
}

func init() {
	godotenv.Load()
	G_CONF = newConf()
	initMySQL()
}

func newConf() *Conf {
	host := os.Getenv("MYSQL_HOST")
	pwd := os.Getenv("MYSQL_PWD")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")
	user := os.Getenv("MYSQL_USER")
	webPW := os.Getenv("WEB_PASSWORD")
	webAddr := os.Getenv("WEB_ADDR")
	secret := os.Getenv("WEB_SECRET")
	domain := os.Getenv("WEB_DOMAIN")
	staticAddr := os.Getenv("WEB_STATIC_ADDR")
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "3306"
	}

	if webAddr == "" {
		webAddr = "127.0.0.1:8948"
	}
	return &Conf{
		MSHost:          host,
		MSPWD:           pwd,
		MSPort:          port,
		MSDBName:        dbName,
		MSUser:          user,
		WebsitePassWord: webPW,
		WebAddr:         webAddr,
		Secret:          secret,
		Domain:          domain,
		StaticAddr:      staticAddr,
	}
}
