package process2

import (
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"day_daily/day_proficient/day_9/chatroom/server/model"
	"day_daily/day_proficient/day_9/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段表示该conn是哪个用户的
	UserId int
}

//通知所有在线的用户的方法
//userId通知其他人上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历上线的人
	userMap := userMgr.GetAllOnlineUser()
	for id, up := range userMap {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesTpye

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json Marshal err: ", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err: ", err)
		return
	}

	//发送，创建tranfer实例发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyOnline err: ", err)
		return
	}
	return


}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error){
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarshal fail err: ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//3.再声明LoginResMes
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	resMes.Data = string(data)

	//6.将resMes 序列化，并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	//7.发送data，将其封装到write函数里
	//因为使用了分层模式（mvc），我们先创建一个transfer实例，然后来读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//编写一个函数serverProcessLogin函数，专门处理登录信息
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	fmt.Println("ServerProcessLogin")
	//1.从message取出mes.data，并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal fail err: ", err)
		return
	}

	//2.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//3.再声明LoginResMes
	var loginResMes message.LoginResMes

	////如果用户id 是100，密码是123456 是合法，否则不合法
	//fmt.Println("判断用户名和密码是否正确")
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//
	//} else {
	//	//不合法
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册再使用..."
	//}

	//我们需要到redis数据库中去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
		////先测试
		//loginResMes.Code = 500
		//loginResMes.Error = "该用户不存在，请注册再使用..."
	} else {
		loginResMes.Code = 200
		//因为用户登录成功，我们就把该登录成功的用户放入到UserMgr里
		this.UserId = loginMes.UserId
		userMgr.AddOnlinerUser(this)
		//通知其他在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id放入到loginResMes.UsersId
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")

	}


	//4.将logResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	//5.将data赋值给resMes
	resMes.Data = string(data)

	//6.将resMes 序列化，并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	//7.发送data，将其封装到write函数里
	//因为使用了分层模式（mvc），我们先创建一个transfer实例，然后来读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
