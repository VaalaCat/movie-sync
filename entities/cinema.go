package entities

import "github.com/sirupsen/logrus"

type Cinema interface {
	GetRoom(name string) (Room, bool)
	GetUserRoom(name string) (Room, bool)
	SetRoom(name string, room Room)
	SetUserRoom(name string, room Room)
	DeleteUserRoom(name string)
	RemoveRoom(name string)
}

type CinemaImpl struct {
	rooms     map[string]Room
	userRooms map[string]Room
}

var (
	c Cinema = (*CinemaImpl)(nil)
)

func init() {
	c = NewCinema()
}

func GetCinema() Cinema {
	return c
}

func NewCinema() *CinemaImpl {
	return &CinemaImpl{
		rooms:     make(map[string]Room),
		userRooms: make(map[string]Room),
	}
}

func (c *CinemaImpl) GetRoom(name string) (Room, bool) {
	r, ok := c.rooms[name]
	return r, ok
}

func (c *CinemaImpl) GetUserRoom(name string) (Room, bool) {
	r, ok := c.userRooms[name]
	return r, ok
}

func (c *CinemaImpl) SetRoom(name string, room Room) {
	c.rooms[name] = room
}

func (c *CinemaImpl) SetUserRoom(name string, room Room) {
	c.userRooms[name] = room
}

func (c *CinemaImpl) RemoveRoom(name string) {
	logrus.Infof("remove room: %s", name)
	delete(c.rooms, name)
	for _, r := range c.userRooms {
		if r.Name() == name {
			delete(c.userRooms, name)
		}
	}
}

func (c *CinemaImpl) DeleteUserRoom(name string) {
	delete(c.userRooms, name)
}
