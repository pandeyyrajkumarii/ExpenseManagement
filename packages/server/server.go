package server

import (
	"ExpenseManagement/internal/userService"
	"github.com/go-chi/chi/v5"
	"k8s.io/klog/v2"
	"net/http"
)

type Server struct {
	Port       string
	UserServer *userService.Server
}

func NewServer(opt ...Opt) *Server {
	svr := &Server{}
	for _, o := range opt {
		o(svr)
	}
	return svr
}

func (s *Server) Start() {
	mux := chi.NewRouter()
	s.registerRoutes(mux)
	if err := http.ListenAndServe(":"+s.Port, mux); err != nil {
		klog.Fatal(err)
	}
	klog.Info("Server started at ", s)
}
