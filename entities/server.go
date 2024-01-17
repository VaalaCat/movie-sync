package entities

import (
	"github.com/gin-gonic/gin"
	socket "github.com/zishang520/socket.io/v2/socket"
)

var server *socket.Server
var router *gin.Engine

func init() {
	server = socket.NewServer(nil, nil)
	router = gin.New()
}

func GetServer() *socket.Server {
	return server
}

func GetRouter() *gin.Engine {
	return router
}
