// +build wireinject

package main

import (
	"github.com/luofeng1/Go-001/Week04/internal/biz"
	"github.com/luofeng1/Go-001/Week04/internal/data"

	"github.com/google/wire"
)

// InitUser ..
func InitUser() *biz.User {
	wire.Build(biz.NewUserDo, data.NewUserRepo)
	return &biz.User{}
}
