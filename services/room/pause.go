package room

import (
	"movie-sync-server/entities"

	"github.com/zishang520/socket.io/v2/socket"
)

func PauseEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, showName := cliMsg.Room, cliMsg.UserName
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		u := r.GetUser(string(client.Id()))
		if u != nil {
			u.SetPlaying(false)
		}
		r.Broadcast("pause", entities.ServerMessage{
			ActionFrom:    "client",
			ActionEmitter: showName,
		})
	}
	return nil
}
