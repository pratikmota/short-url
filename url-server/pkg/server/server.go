package server

import (
	"url-server/pkg/http_multiplexer"
)

// Server ...
type Server struct {
	mux http_multiplexer.IMultiplexer
}

// New creates a new *Server
func New(mux http_multiplexer.IMultiplexer) *Server {
	return &Server{mux: mux}
}

// Serve allow the server to start serving
func (s Server) Serve() {
	s.mux.Serve()
}
