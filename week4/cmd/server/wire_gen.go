// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"week4/internal/biz"
	"week4/internal/data"
)

// Injectors from wire.go:

func InitUserUsecase() *biz.UserUsecase {
	userRepo := data.NewUserRepo()
	userUsecase := biz.NewUserUsecase(userRepo)
	return userUsecase
}