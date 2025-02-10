package userService

import (
	"ExpenseManagement/internal/userService/contracts"
	"ExpenseManagement/internal/userService/service"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonOut, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonOut)

	klog.Infof("Create User success", user)
}
