package model

import (
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//我们在服务器启动后，就初始化一个userDao实例，
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var(
	MyUserDao *UserDao
)

//定义一个userDAO结构体
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool : pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error){
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里我们反序列化user实例
	//获取的res数据接口应该是："{\"userId\":100,\"userPwd\":\"123456\",\"userName\":\"\"}"
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("userDao json umarshal err: ", err)
		return
	}
	return
}

//完成登录后校验
func (this *UserDao) Login (userId int, userPwd string) (user *User, err error) {
	//先从userDao连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这个用户获取到了，但是密码不一定正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register (user *message.User) (err error) {
	//先从userDao连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err != nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户处错误 err：", err)
		return
	}
	return
}
