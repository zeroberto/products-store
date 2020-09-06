package cmd

import "log"

// Server is responsible for starting application services
type Server struct{}

// Start is responsible for initializing the server settings and routes
func (s *Server) Start() {
	log.Println("Starting server...")
}

// Stop is responsible for shutting down the server
func (s *Server) Stop() {
}
