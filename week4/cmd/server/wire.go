//+build wireinject 忽略编译

package main

import (
	"week4/internal/biz"
	"week4/internal/data"

	"github.com/google/wire"
)

func InitUserUsecase() *biz.UserUsecase {
	wire.Build(biz.NewUserUsecase, data.NewUserRepo)
	return &biz.UserUsecase{}
}