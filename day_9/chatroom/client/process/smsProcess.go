package process

import (
	"day_daily/day_proficient/day_9/chatroom/client/utils"
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个message
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail: ", err)
		return
	}
	mes.Data = string(data)

	//4.再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail: ", err)
		return
	}

	//5.将序列化之后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupMes err: ", err)
		return
	}
	return
}
