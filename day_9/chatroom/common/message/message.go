package message

const (
	LoginMesType  = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesTpye = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusystatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}


type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`
	UsersId []int //保存用户id的切片
	Error string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"`//返回状态码 400表示该用户已经占用，200表示注册成功
	Error string  `json :"error"` //返回错误信息
}

//定义通知的消息类型，用于用户状态变化
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"` //用户状态
}

//增加一个SmsMes 发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User //匿名结构体，
}