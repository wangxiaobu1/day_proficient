package process2

import (
	"day_daily/day_proficient/day_9/chatroom/client/utils"
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//取出message的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("jsom umarshl err: ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err: ", err)
		return
	}

	//遍历服务器端的onlineUsers，并将消息转发
	for id, up := range userMgr.onlineUsers {
		//过滤自己，不要自己发自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个tranfer实例，发送data
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err: ", err)
	}
}
