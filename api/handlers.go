package api

import "net/http"

func (s *Server) HandleFoo(w http.ResponseWriter, r *http.Request) {}
func (s *Server) HandleBar(w http.ResponseWriter, r *http.Request) {}
func (s *Server) HandleBaz(w http.ResponseWriter, r *http.Request) {}
