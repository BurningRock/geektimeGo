package service

import (
	v1 "week4/api/user/v1"
	"week4/internal/biz"
	"context"
)

type UserService struct {
	u *biz.UserUsecase
	v1.UnimplementedUserServer
}

func NewUserService(u *biz.UserUsecase) v1.UserServer {
	return &UserService{u: u}
}

func (s *UserService) RegisterUser(ctx context.Context, r *v1.RegisterUserRequest) (*v1.RegisterUserReply, error) {
	// dto -> do
	u := &biz.User{Name: r.Name, Age: r.Age}

	// call biz
	s.u.SaveUser(u)

	// return reply
	return &v1.RegisterUserReply{Id: u.ID}, nil
}