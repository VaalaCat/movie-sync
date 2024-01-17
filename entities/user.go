package entities

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zishang520/socket.io/v2/socket"
)

type User interface {
	ID() string
	Send(event string, message ServerMessage)
	GetSocket() *socket.Socket
	GetTime() int64
	SetTime(time int64)
	SetID(id string)
	SetSocket(socket *socket.Socket)
	SetUsername(username string)
	GetUserName() string
	IsPlaying() bool
	SetPlaying(playing bool)
}

type UserImpl struct {
	id         string
	playTime   int64
	username   string
	socket     *socket.Socket
	updateTime int64
	playing    bool
}

var (
	_ User = (*UserImpl)(nil)
)

func (u *UserImpl) ID() string {
	return u.id
}

func (u *UserImpl) Send(event string, message ServerMessage) {
	if u == nil {
		return
	}
	rawMsg, err := json.Marshal(message)
	if err != nil {
		logrus.WithError(err).Errorf("send event to user marshal json [%v]:{%v} error", event, message)
	}
	u.socket.Emit(event, string(rawMsg))
}

func (u *UserImpl) SetTime(playTime int64) {
	if u == nil {
		return
	}
	u.updateTime = time.Now().Unix()
	u.playTime = playTime
}

func (u *UserImpl) GetTime() int64 {
	if u == nil {
		return 0
	}
	return u.playTime
}

func (u *UserImpl) GetSocket() *socket.Socket {
	if u == nil {
		return nil
	}
	return u.socket
}

func (u *UserImpl) SetID(id string) {
	if u == nil {
		return
	}
	u.updateTime = time.Now().Unix()
	u.id = id
}

func (u *UserImpl) SetSocket(socket *socket.Socket) {
	if u == nil {
		return
	}
	u.socket = socket
}

func (u *UserImpl) SetUsername(username string) {
	if u == nil {
		return
	}
	u.username = username
}

func (u *UserImpl) GetUserName() string {
	if u == nil {
		return ""
	}
	return u.username
}

func (u *UserImpl) IsPlaying() bool {
	if u == nil {
		return false
	}
	return u.playing
}

func (u *UserImpl) SetPlaying(playing bool) {
	if u == nil {
		return
	}
	u.playing = playing
}
