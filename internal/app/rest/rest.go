package rest

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/deck/internal/pkg/deck"
)

type Server struct {
	port        string
	logger      *slog.Logger
	router      *gin.Engine
	deckManager deck.Manager
}

type API interface {
	Serve()
}

func New(port string, logger *slog.Logger, deckManager deck.Manager) *Server {
	s := &Server{
		router:      gin.Default(),
		port:        port,
		logger:      logger,
		deckManager: deckManager,
	}

	s.setupConfig(s.router)
	s.setupEndpoint()
	return s
}

func (s *Server) Serve() {
	s.router.Run("0.0.0.0:" + s.port)
}

func (s *Server) setupConfig(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
}
