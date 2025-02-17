package service

import (
	"ExpenseManagement/internal/userService/contracts"
	"ExpenseManagement/internal/userService/repo"
	"ExpenseManagement/packages/middleware"
	"fmt"
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
	err := s.ValidateCreateUserRequest(*req)
	if err != nil {
		klog.Infof("validation faliure", "error", err.Error())
		return nil, err
	}
	user, err := s.Repo.FindByUserId(req.Id)
	if err != nil && err.Error() != "user_not_found" {
		klog.Infof("DB error")
		return nil, err
	}
	if user != nil {
		klog.Infof("User already exists")
		return nil, fmt.Errorf("user_already_exists")
	}

	userDB, err := s.Repo.Create(req)
	if err != nil {
		return nil, err
	}
	klog.Infof(" User created success", "user:", userDB)

	return &contracts.UserResponse{UserId: userDB.Id}, nil
}

func (s *Service) LoginUser(req *contracts.UserLogin) (*contracts.LoginResponse, error) {
	//validate
	err := s.ValidateUserLoginRequest(*req)
	if err != nil {
		klog.Infof("validation faliure", "error", err.Error())
		return nil, err
	}

	//Find User By Id
	user, err := s.Repo.FindUserByIdForLogin(req.Id)
	if err != nil && err.Error() != "user_not_found" {
		klog.Infof("DB error")
		return nil, err
	}
	if user == nil {
		klog.Infof("User does not exist")
		return nil, fmt.Errorf("user_not_found")
	}

	//check password
	if user.Password != req.Password {
		klog.Infof("incorect_password")
		return nil, fmt.Errorf("either_incorect_user_or_password")
	}
	// Generate JWT token
	token, err := middleware.GenerateJWT(user.UserId)
	if err != nil {
		klog.Infof("JWT generation failed")
		return nil, err
	}

	klog.Infof(" User logged in successfully", "user:", user.UserId)

	return &contracts.LoginResponse{Token: token}, nil
}
