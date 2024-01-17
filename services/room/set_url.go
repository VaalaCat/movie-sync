package room

import (
	"movie-sync-server/entities"

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
	}
	return nil
}
