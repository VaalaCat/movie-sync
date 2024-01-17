package entities

type ClientMessage struct {
	UserName string `json:"username,omitempty"`
	Room     string `json:"room,omitempty"`
	URL      string `json:"url,omitempty"`
	Time     int64  `json:"time,omitempty"`
	Playing  bool   `json:"playing,omitempty"`
	Payload  string `json:"payload,omitempty"`
}

type ServerMessage struct {
	URL           string       `json:"url,omitempty"`
	UserStatus    []UserStatus `json:"userStatus,omitempty"`    // timepair[0] = min, timepair[1] = max
	ActionFrom    string       `json:"actionFrom,omitempty"`    // "server" or "client"
	ActionEmitter string       `json:"actionEmitter,omitempty"` // username
}

type UserStatus struct {
	UserID   string `json:"userID,omitempty"`
	UserName string `json:"username,omitempty"`
	Time     int64  `json:"time,omitempty"`
	Playing  bool   `json:"playing,omitempty"`
}
