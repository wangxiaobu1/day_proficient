package model

const (
	UserStatusOnline = 1
	UserStatusOffine = iota
	)

type User struct {
	UserId int `json:"user_id"`
	Passwd string `json:"passwd"`
	Nick string `jsong:"nick"`
	Sex string `json:"sex"`
	Header string `json:"header"`
	LastLogin string `json:"last_login"`
	Status string `json:"status"`
}
