package process

import (
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	//显示即可
	//1.反序列化你
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), smsMes)
	if err != nil {
		fmt.Println("json unmarshal err: ", err.Error())
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id：\t%d 对大家说：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
