package service

import (
	"github.com/luofeng1/Go-000/Week02/homework/dao"
	"github.com/pkg/errors"
)

// GetUser 获取用户信息
func GetUser(userID string) (*dao.User, error) {
	user, err := dao.GetByID(userID)
	return user, errors.WithMessage(err, "service getUser")
}
