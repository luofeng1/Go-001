package dao

import (
	"database/sql"
	"errors"

	"github.com/luofeng1/Go-000/Week02/homework/code"
)

// User 用户信息
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetByID 通过用户id查询用户
func GetByID(id string) (*User, error) {
	u := &User{}
	result := db.Model(u).Where("id = ?", id).First(u)
	if result.Error == nil {
		return u, nil
	}
	if errors.Is(result.Error, sql.ErrNoRows) {
		return nil, code.UserNotExists
	}
	return nil, code.DBUnKnowErr
}
