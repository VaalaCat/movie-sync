package services

import (
	"encoding/json"
	"fmt"
	"movie-sync-server/conf"
	"movie-sync-server/entities"
	"movie-sync-server/services/room"
	"movie-sync-server/utils"
	"net/http"

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
	// //如果有一个客户端连接上
	// server.OnConnect("/", room.ConnectEndpoint)
	// server.OnEvent("/", "join", room.JoinEndpoint)
	// server.OnEvent("/", "setUrl", room.SetUrlEndpoint)
	// server.OnEvent("/", "time", room.TimeEndpoint)
	// // 如果接收到sync请求则发送时间的最大最小值
	// server.OnEvent("/", "sync", room.SyncEndpoint)
	// // 如果有用户发出getTime请求则开始广播getTime请求
	// server.OnEvent("/", "getTime", room.GetTimeEndpoint)
	// // 如果有用户发出setTime请求则使用setTime消息将所有用户的时间同步
	// server.OnEvent("/", "setTime", room.SetTimeEndpoint)
	// // 用户发出getUsers请求则返回所有的用户名
	// server.OnEvent("/", "getUsers", room.GetUsersEndpoint)
	// // 用户发出getUrl请求则返回url给对应用户
	// server.OnEvent("/", "getUrl", room.GetUrlEndpoint)
	// // 用户发出play则广播play请求
	// server.OnEvent("/", "play", room.PlayEndpoint)
	// // 用户发出stop则广播stop请求
	// server.OnEvent("/", "pause", room.PauseEndpoint)
	// server.OnError("/", room.ErrorEndpoint)
	// server.OnDisconnect("/", room.DisconnectEndpoint)

	// go func() {
	// 	if err := server.Serve(); err != nil {
	// 		logrus.Fatalf("socketio listen error: %s\n", err)
	// 	}
	// }()
}

func RouterHandler() {
	server := entities.GetServer()
	router := entities.GetRouter()

	router.Use(utils.CORSMiddleware())
	router.GET("/socket.io/*any", gin.WrapH(server.ServeHandler(nil)))
	router.POST("/socket.io/*any", gin.WrapH(server.ServeHandler(nil)))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/movie")
	})
	router.StaticFile("/movie", "../asset/index.html")
	router.StaticFile("/movie/login", "../asset/index.html")
	router.StaticFS("/movie/css", http.Dir("../asset/css"))
	router.StaticFS("/movie/js", http.Dir("../asset/js"))
	router.NoRoute(func(c *gin.Context) {
		c.File("../asset/index.html")
	})
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
