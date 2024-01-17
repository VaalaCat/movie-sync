package room

import (
	"movie-sync-server/entities"

	"github.com/zishang520/socket.io/v2/socket"
)

func SetTimeEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, showName, time := cliMsg.Room, cliMsg.UserName, cliMsg.Time
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		r.Broadcast("setTime", entities.ServerMessage{
			ActionFrom:    "client",
			ActionEmitter: showName,
			UserStatus:    r.GetAllUserStatus(),
		})
		r.Broadcast("pause", entities.ServerMessage{
			ActionFrom:    "client",
			ActionEmitter: showName,
		})
		for _, u := range r.GetUsers() {
			u.SetTime(time)
		}
	}
	return nil
}
