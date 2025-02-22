package txnService

import (
	"ExpenseManagement/internal/txnService/contracts"
	"ExpenseManagement/internal/txnService/service"
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

func (s *Server) SaveTransaction(w http.ResponseWriter, r *http.Request) {
	var txn contracts.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User not found, sign in", http.StatusUnauthorized)
		return
	}
	klog.Infof("Saving transaction: %v", txn)
	resp, err := s.Service.SaveTransaction(&txn, userID)
	if err != nil {
		utils.HandleErrorResponse(w, err)
		return
	}
	utils.HandleResponse(w, resp)
	klog.Infof("CreateTxn Response: %v", resp)
}

func (s *Server) GetTransactionByFilter(w http.ResponseWriter, r *http.Request) {
	var req contracts.GetTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User not found, sign in", http.StatusUnauthorized)
		return
	}
	klog.Infof("Saving transaction: %v", req)
	resp, err := s.Service.GetTransactionByFilter(&req, userID)
	if err != nil {
		utils.HandleErrorResponse(w, err)
		return
	}
	utils.HandleResponse(w, resp)
	klog.Infof("CreateTxn Response: %v", resp)
}
