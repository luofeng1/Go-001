// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/luofeng1/Go-001/Week04/internal/biz"
	"github.com/luofeng1/Go-001/Week04/internal/data"
)

// Injectors from wire.go:

func InitUser() *biz.User {
	userRepo := data.NewUserRepo()
	user := biz.NewUserDo(userRepo)
	return user
}