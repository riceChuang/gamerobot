package model

type GameInfo struct {
	GameID  int32
	Name    string
	Buttons map[string]*Message
}
