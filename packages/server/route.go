package server

import (
	"ExpenseManagement/packages/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) registerRoutes(mux *chi.Mux) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to your space..."))
		w.WriteHeader(http.StatusOK)
	})

	mux.Post("/expense/signup", s.UserServer.CreateUser)
	mux.Post("/expense/login", s.UserServer.LoginUser)
	mux.With(middleware.AuthMiddleware).Post("/expense/transaction", s.TxnServer.SaveTransaction)
	mux.With(middleware.AuthMiddleware).Post("/expense/transaction/fetch", s.TxnServer.GetTransactionByFilter)
}
