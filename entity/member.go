package entity

import "time"

type Member struct {
	MemberId int64
	Name     string
	Devices  []*Device
}

type Message struct {
	Version   int
	MessageId int64
	Image     []byte
	Text      string
	From      int64 // Member
	To        int64 // Device
	Date      time.Time
}

type ChatRoom struct {
	RoomId  int64
	Members []*Member
}

type Device struct {
	DeviceId int64
	DevType  int
}
