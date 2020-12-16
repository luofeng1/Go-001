package data

import (
	"fmt"

	"github.com/luofeng1/Go-001/Week04/internal/biz"
)

// check userRepo 实现
var _ biz.UserRepo = new(userRepo)

// NewUserRepo 实现持久化user对象
func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

// userRepo ..
type userRepo struct{}

// Save ..
func (u *userRepo) Save(user *biz.UserPo) int32 {
	fmt.Printf("SQL: INSERT INTO user (name, age) VALUES (%s, %d);\n", user.Name, user.Age)
	// TODO: 执行数据库,缓存操作,保存并返回用户id
	return 1
}
