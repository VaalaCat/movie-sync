package room

import (
	"movie-sync-server/entities"

	"github.com/zishang520/socket.io/v2/socket"
)

func UpdateInfoEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, time := cliMsg.Room, cliMsg.Time
	username := string(client.Id())
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		u := r.GetUser(username)
		if u != nil {
			u.SetTime(time)
			u.SetPlaying(cliMsg.Playing)
		}
	}
	return nil
}
