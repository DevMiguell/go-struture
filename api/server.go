package api

import (
	"encoding/json"
	"net/http"
	"v1/storage"
	"v1/util"
)

type Server struct {
	linstenAddr string
	store       storage.Storage
}

func NewServer(linstenAddr string, store storage.Storage) *Server {
	return &Server{
		linstenAddr: linstenAddr,
		store:       store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/user/id", s.handleDeleteUserByID)

	return http.ListenAndServe(s.linstenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	_ = util.Round2Dec(10.34434)

	json.NewEncoder(w).Encode(user)
}
