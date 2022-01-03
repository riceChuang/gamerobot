package model

type GameInfo struct {
	GameID   int32
	Name     string
	RoomType []string
	Buttons  map[string]*Message
}
