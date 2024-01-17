package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Setting struct {
	Port        string
	Addr        string
	AllowOrigin string
	YoukuCcode  string
	YoukuCkey   string
}

var ServerSetting Setting

func init() {
	Init()
}

func Init() {
	// load from .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	allowOrigin := os.Getenv("ALLOW_ORIGIN")
	if port == "" {
		port = "9999"
	}
	if allowOrigin == "" {
		allowOrigin = "*"
	}

	youkuCcode := os.Getenv("YOUKU_CCODE")
	youkuCkey := os.Getenv("YOUKU_CKEY")

	ServerSetting = Setting{
		Port:        port,
		AllowOrigin: allowOrigin,
		YoukuCcode:  youkuCcode,
		YoukuCkey:   youkuCkey,
	}
}
