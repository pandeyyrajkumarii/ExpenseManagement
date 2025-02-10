package service

import (
	"ExpenseManagement/internal/userService/contracts"
	"ExpenseManagement/internal/userService/repo"
	"k8s.io/klog/v2"
)

type Service struct {
	Repo *repo.UserRepo
}

func NewService(opt ...Opt) *Service {
	sv := &Service{}
	for _, o := range opt {
		o(sv)
	}
	return sv
}

func (s *Service) CreateUser(req *contracts.User) (*contracts.UserResponse, error) {
	// add validation logic
	userDB, err := s.Repo.Create(req)
	if err != nil {
		return nil, err
	}
	klog.Infof(" User created success", "user:", userDB)

	return &contracts.UserResponse{UserId: userDB.Id}, nil
}
