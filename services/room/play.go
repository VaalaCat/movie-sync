package room

import (
	"movie-sync-server/entities"

	"github.com/zishang520/socket.io/v2/socket"
)

func PlayEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, showName := cliMsg.Room, cliMsg.UserName
	userID := string(client.Id())
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		u := r.GetUser(userID)
		if u != nil {
			u.SetPlaying(true)
		}
		r.Broadcast("play", entities.ServerMessage{
			ActionFrom:    "client",
			ActionEmitter: showName,
		})
	}
	return nil
}
