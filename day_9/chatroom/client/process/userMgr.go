package process

import (
	"day_daily/day_proficient/day_9/chatroom/client/model"
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"fmt"
)

//客户端维护的map
var onlineUser map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用户登录成功后，完成对CurUser的初始化

//客户端显示当前在线人数
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUser {
		fmt.Println("用户ID：\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUser[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUser[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
