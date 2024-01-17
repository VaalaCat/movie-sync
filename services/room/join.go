package room

import (
	"movie-sync-server/entities"

	"github.com/sirupsen/logrus"
	"github.com/zishang520/socket.io/v2/socket"
)

func JoinEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	roomname, showName := cliMsg.Room, cliMsg.UserName
	username := string(client.Id())
	if roomname == "" || username == "" || showName == "" {
		logrus.Warnf("room or username is empty: room: [%s], username: [%s], showName: [%s]", roomname, username, showName)
		return nil
	}
	var newUser entities.User = &entities.UserImpl{}

	//首先判断当前用户是否想要加入已有的房间，如果房间不存在则新建房间
	joined := false
	var joinedRoom entities.Room
	if r, ok := entities.GetCinema().GetRoom(roomname); ok {
		u := r.GetUser(username)
		if u != nil {
			logrus.Warnf("user [%s] already in room [%s]", username, roomname)
			return nil
		}
		joinedRoom = r
		newUser.SetID(username)
		newUser.SetSocket(client)
		newUser.SetUsername(showName)
		r.AddUser(newUser)
		joined = true
	}
	if !joined {
		var newRoom entities.Room = &entities.RoomImpl{}
		newRoom.SetName(roomname)
		newRoom.InitUsers()
		newUser.SetID(username)
		newUser.SetUsername(showName)
		newUser.SetSocket(client)
		newRoom.AddUser(newUser)
		joinedRoom = newRoom
		entities.GetCinema().SetRoom(roomname, newRoom)
	}

	joinedRoom.Broadcast("userjoin", entities.ServerMessage{ActionFrom: "client", ActionEmitter: showName})
	joinedRoom.GetUsers()
	newUser.Send("rootinit", entities.ServerMessage{
		ActionFrom:    "server",
		URL:           joinedRoom.GetUrl(),
		UserStatus:    joinedRoom.GetAllUserStatus(),
		ActionEmitter: "server",
	})

	logrus.Infof("user [%s] join room [%s] success", username, roomname)
	return nil
}
