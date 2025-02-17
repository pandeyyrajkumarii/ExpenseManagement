package userService

import (
	"ExpenseManagement/internal/userService/contracts"
	"ExpenseManagement/internal/userService/service"
	"ExpenseManagement/internal/utils"
	"encoding/json"
	"k8s.io/klog/v2"
	"net/http"
)

type Server struct {
	Service *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{Service: svc}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user contracts.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := s.Service.CreateUser(&user)
	if err != nil {
		utils.HandleErrorResponse(w, err)
		return
	}
	utils.HandleResponse(w, resp)

	klog.Infof("Create User success", user)
}

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	var login contracts.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := s.Service.LoginUser(&login)
	if err != nil {
		utils.HandleErrorResponse(w, err)
		return
	}
	utils.HandleResponse(w, resp)

	klog.Infof("logged in successfully!!!", login)
}
