package config

type Server struct {
	address string
	port    uint
}

func NewServer() *Server {
	return &Server{
		"localhost",
		1984,
	}
}
