package data


import (
	"log"
	"week4/internal/biz"
)
var _ biz.UserRepo = new(userRepo) //查看data层是否实现某个接口 依赖倒置 biz层实现某个接口 data层实现某个实现

const (
	userID = 100
)

func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

type userRepo struct{}

func (r *userRepo) Save(u *biz.User) int32 {
	log.Printf("save user, name: %s, age: %d, id: %d", u.Name, u.Age, userID)  //模拟出某个对象 其实只是为了获取该ID
	return userID
}