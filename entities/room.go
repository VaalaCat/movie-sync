package entities

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	socket "github.com/zishang520/socket.io/v2/socket"
)

type Room interface {
	Name() string
	SetName(name string)
	SetUrl(url string)
	GetUrl() string
	GetMinTime() int64
	GetMaxTime() int64
	AddUser(user User)
	RemoveUser(username string)
	Broadcast(event string, message ServerMessage)
	GetUsers() []User
	GetUser(name string) User
	GetAllUserStatus() []UserStatus
	InitUsers()
	Refresh()
}

var (
	_ Room = (*RoomImpl)(nil)
)

type RoomImpl struct {
	name        string
	url         string
	users       map[string]User
	lastPlay    time.Time
	lastStop    time.Time
	setTimeLock sync.Mutex
}

func (r *RoomImpl) Name() string {
	return r.name
}

func (r *RoomImpl) SetName(name string) {
	r.name = name
	r.lastPlay = time.Now()
	r.lastStop = time.Now()
}

func (r *RoomImpl) SetUrl(url string) {
	r.url = url
}

func (r *RoomImpl) GetUrl() string {
	return r.url
}

func (r *RoomImpl) GetMinTime() int64 {
	var times []int64
	for _, user := range r.users {
		times = append(times, user.GetTime())
	}
	min := times[0]
	for _, time := range times {
		if time < min {
			min = time
		}
	}
	return min
}

func (r *RoomImpl) GetMaxTime() int64 {
	var times []int64
	for _, user := range r.users {
		times = append(times, user.GetTime())
	}
	max := times[0]
	for _, time := range times {
		if time > max {
			max = time
		}
	}
	return max
}

func (r *RoomImpl) AddUser(user User) {
	r.users[user.ID()] = user
	GetCinema().SetUserRoom(user.ID(), r)
	user.GetSocket().Join(socket.Room(r.name))
}

func (r *RoomImpl) RemoveUser(username string) {
	tmpUser, ok := r.users[username]
	if ok {
		tmpUser.GetSocket().Leave(socket.Room(r.name))
		GetCinema().DeleteUserRoom(tmpUser.ID())
		delete(r.users, username)
	}
}

func (r *RoomImpl) Broadcast(event string, message ServerMessage) {
	rawMsg, err := json.Marshal(message)
	if err != nil {
		logrus.WithError(err).Errorf("broadcast event marshal json [%v]:{%v} error", event, message)
	}

	if event == "play" {
		r.lastPlay = time.Now()
	}
	if event == "pause" {
		r.lastStop = time.Now()
	}
	if event == "setTime" {
		if r.setTimeLock.TryLock() {
			logrus.Infof("set time lock")
			r.setTimeLock.Lock()
			go func() {
				time.Sleep(time.Millisecond * 10)
				r.setTimeLock.Unlock()
			}()
		} else {
			logrus.Infof("set time already locked")
			return
		}
	}
	err = GetServer().To(socket.Room(r.name)).Emit(event, string(rawMsg))
	if err != nil {
		logrus.WithError(err).Errorf("broadcast event [%v]:{%v} error", event, message)
	}
}

func (r *RoomImpl) InitUsers() {
	r.users = make(map[string]User)
}

func (r *RoomImpl) Refresh() {
	GetServer().To(socket.Room(r.name)).Emit("refresh")
}

func (r *RoomImpl) GetUsers() []User {
	return lo.MapToSlice(r.users, func(k string, v User) User { return v })
}

func (r *RoomImpl) GetUser(name string) User {
	user, ok := r.users[name]
	if ok {
		return user
	}
	return nil
}

func (r *RoomImpl) GetAllUserStatus() []UserStatus {
	return lo.MapToSlice(r.users, func(k string, v User) UserStatus {
		return UserStatus{
			UserName: v.GetUserName(),
			UserID:   v.ID(),
			Time:     v.GetTime(),
			Playing:  v.IsPlaying(),
		}
	})
}
