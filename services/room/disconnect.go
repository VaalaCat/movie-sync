package room

import (
	"movie-sync-server/entities"
	"time"

	"github.com/zishang520/socket.io/v2/socket"
)

func DisconnectEndpoint(client *socket.Socket) []byte {
	userID := string(client.Id())
	if r, ok := entities.GetCinema().GetUserRoom(userID); ok {
		r.RemoveUser(userID)
		if len(r.GetUsers()) == 0 {
			roomName := r.Name()
			go func() {
				time.Sleep(60 * time.Second)
				if r, ok := entities.GetCinema().GetRoom(roomName); ok {
					if len(r.GetUsers()) == 0 {
						entities.GetCinema().RemoveRoom(roomName)
					}
				}
			}()
		} else {
			r.Broadcast("leaveRoom", entities.ServerMessage{
				ActionFrom:    "server",
				UserStatus:    r.GetAllUserStatus(),
				ActionEmitter: string(client.Id()),
			})
		}
	}
	return nil
}
