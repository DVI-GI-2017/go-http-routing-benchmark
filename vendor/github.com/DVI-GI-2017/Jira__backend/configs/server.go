package configs

import (
	"os"
	"strconv"
)

type Server struct {
	Port int `json:"port"`
}

func (s *Server) GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(s.Port)
	}

	return port
}
