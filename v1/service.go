package services

import (
	"encoding/json"
	"jjavilara17/rcpt/models"
	"net/http"

	"github.com/gorilla/mux"
)

type Serv struct{}

func (s *Serv) ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	json.NewEncoder(w).Encode(response)
}

func (s *Serv) GetPoints(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	json.NewEncoder(w).Encode(response)
}

func (s *Serv) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("path", s.ProcessReceipts).Methods(http.MethodPost)
	r.HandleFunc("path2", s.GetPoints).Methods(http.MethodGet)
}
