package service

import (
	"context"

	v1 "github.com/luofeng1/Go-001/Week04/api/user/v1"
	"github.com/luofeng1/Go-001/Week04/internal/biz"
)

// UserService ..
type UserService struct {
	// 领域对象
	u *biz.User
	v1.UnimplementedUserServer
}

// NewUserService ..
func NewUserService(u *biz.User) v1.UserServer {
	return &UserService{u: u}
}

// RegisterUser ..
func (s *UserService) RegisterUser(ctx context.Context, r *v1.RegisterUserRequest) (*v1.RegisterUserReply, error) {
	// DTO -> DO
	u := &biz.UserPo{Name: r.Name, Age: r.Age}

	// 调用领域的方法save
	s.u.Save(u)

	// return reply
	return &v1.RegisterUserReply{Id: u.ID}, nil
}
