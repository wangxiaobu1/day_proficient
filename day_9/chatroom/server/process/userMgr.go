package process2

import "fmt"

//UserMgr实例，在服务器端有且仅有一个
//因为在很多地方都使用，因此是全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers的添加
func (this *UserMgr) AddOnlinerUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//完成对onlineUsers删除
func (this *UserMgr) DelOnlinerUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		//说明要查找的userId当前不在线
		err = fmt.Errorf("用户%id 不存在", userId)
		return
	}
	return
}
