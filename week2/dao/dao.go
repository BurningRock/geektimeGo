package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"fmt"
)

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答：应该将sql.ErrNoRows抛给上层，同时有一定的上下文信息
type User struct {
	Id uint64 `json:"Id"`
	Name string `json: "name""`
}

var db *sql.DB

func GetUserById(id uint64) (user *User, err error) {
	user = &User{Id:id}
	err = db.QueryRow("select * from user where id=?",id).Scan(&user)
	if err !=nil {
		if errors.Is(err,sql.ErrNoRows) {
			return nil, errors.Wrap(err,fmt.Sprintf("find user null,user id: %v",id))
		}else {
			return nil,err
		}
	}
	return
}