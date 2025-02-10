package server

import "ExpenseManagement/internal/userService"

type Opt func(server *Server)

func WithPort(port string) Opt {
	return func(server *Server) {
		server.Port = port
	}
}

func WithUserServer(usrSvr *userService.Server) Opt {
	return func(server *Server) {
		server.UserServer = usrSvr
	}
}
