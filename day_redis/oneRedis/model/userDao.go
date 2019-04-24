package model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"log"
)

//定义一个全局变量
var MyUserDao *UserDao

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool)(userDao *UserDao) {
	userDao = &UserDao{
		pool : pool,
	}
	return
}

//定义一个UserDao结构体
//UserDao操作redis
type UserDao struct {
	pool *redis.Pool
}

//根据用户id，返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//redis里查询用户id
	res, err := redis.String(conn.Do("HGet", "user", id))
	if err != nil {
		//判断是否存在用户
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXITS
		}
		return
	}
	user = &User{}
	//把获取的res反序列化user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		log.Fatalln("解码失败")
	}
	return
}

//完成登录的校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//用户信息获取到了，但不知道是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWDNOTTRUE
	}
	return
}