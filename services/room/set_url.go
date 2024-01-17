package room

import (
	"movie-sync-server/entities"

	"github.com/sirupsen/logrus"
	"github.com/zishang520/socket.io/v2/socket"
)

func SetUrlEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, showName, url := cliMsg.Room, cliMsg.UserName, cliMsg.URL
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		r.SetUrl(url)
		r.Broadcast("setUrl", entities.ServerMessage{
			URL:           url,
			ActionEmitter: showName,
			ActionFrom:    "client",
		})
	} else {
		logrus.Warnf("room [%s] not found, can not set url", room)
	}
	return nil
}
