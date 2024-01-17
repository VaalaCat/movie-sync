package room

import (
	"encoding/json"
	"movie-sync-server/entities"

	"github.com/sirupsen/logrus"
	"github.com/zishang520/socket.io/v2/socket"
)

func GetInfoEndpoint(client *socket.Socket, cliMsg *entities.ClientMessage) []byte {
	room, showName := cliMsg.Room, cliMsg.UserName
	userID := string(client.Id())
	if r, ok := entities.GetCinema().GetRoom(room); ok {
		u := r.GetUser(userID)
		if u != nil {
			u.Send("roomInfo", entities.ServerMessage{
				URL:           r.GetUrl(),
				UserStatus:    r.GetAllUserStatus(),
				ActionFrom:    "client",
				ActionEmitter: showName,
			})
		} else {
			logrus.Warnf("user [%s] not in room [%s]", userID, room)
			rawMsg, err := json.Marshal(entities.ServerMessage{
				URL:           r.GetUrl(),
				UserStatus:    r.GetAllUserStatus(),
				ActionFrom:    "client",
				ActionEmitter: showName,
			})
			if err != nil {
				logrus.WithError(err).Error("json marshal error")
				return nil
			}
			client.Emit("roomInfo", string(rawMsg))
		}
		return nil
	}
	return nil
}
