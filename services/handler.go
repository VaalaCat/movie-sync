package services

import (
	"embed"
	"encoding/json"
	"fmt"
	"movie-sync-server/conf"
	"movie-sync-server/entities"
	"movie-sync-server/services/room"
	"movie-sync-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zishang520/socket.io/v2/socket"
)

func Wrapper(event string, client *socket.Socket, handler func(client *socket.Socket, cliMsg *entities.ClientMessage) []byte) func(...any) {
	return func(datas ...any) {
		logrus.Infof("recive a client event: [%s], datas: %+v", event, datas)
		d, ok := datas[0].(string)
		if !ok || d == "" {
			logrus.Errorf("invalid data type: %+v", datas)
			return
		}

		clientMsg := &entities.ClientMessage{}
		if err := json.Unmarshal([]byte(d), clientMsg); err != nil {
			logrus.WithError(err).Errorf("invalid data type: %+v", datas)
			return
		}

		r := handler(client, clientMsg)
		if len(r) > 0 {
			client.Emit(event)
		}
	}
}

func EventHandler() {
	server := entities.GetServer()
	server.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		//连接后必须加入房间
		client.On("join", Wrapper("join", client, room.JoinEndpoint))
		// 设置房间的视频地址
		client.On("setUrl", Wrapper("setUrl", client, room.SetUrlEndpoint))
		// 如果用户回复了时间则将对应用户时间更新
		client.On("updateMyInfo", Wrapper("updateMyInfo", client, room.UpdateInfoEndpoint))
		// 如果接收到getInfo请求则发送房间所有信息
		client.On("getRoomInfo", Wrapper("getRoomInfo", client, room.GetInfoEndpoint))
		client.On("play", Wrapper("play", client, room.PlayEndpoint))
		client.On("pause", Wrapper("pause", client, room.PauseEndpoint))
		client.On("setTime", Wrapper("setTime", client, room.SetTimeEndpoint))
		client.On("disconnect", func(a ...any) {
			logrus.Infof("client disconnected: %s", client.Id())
			room.DisconnectEndpoint(client)
		})
	})
}

func RouterHandler(fs embed.FS) {
	server := entities.GetServer()
	router := entities.GetRouter()

	router.Use(utils.CORSMiddleware())
	HandleStaticFile(fs, router)
	router.GET("/socket.io/*any", gin.WrapH(server.ServeHandler(nil)))
	router.POST("/socket.io/*any", gin.WrapH(server.ServeHandler(nil)))
}

func Run() {
	router := entities.GetRouter()
	if err := router.Run(fmt.Sprintf(":%s", conf.ServerSetting.Port)); err != nil {
		logrus.Fatal("failed run app: ", err)
	}
	defer entities.GetServer().Close(func(err error) {
		logrus.Fatal(err)
	})
}
