package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gostack-labs/adminx/configs"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
)

// Server serves HTTP requests for our adminx service.
type Server struct {
	config configs.Config
	store  db.Store
	router *bytego.App
}

// NewServer create a new HTTP server and set up routing.
func NewServer(config configs.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := bytego.New()
	router.Debug(true)
	router.Use(logger.New())
	router.Validator(validator.New().Struct)
	router.GET("/", func(c *bytego.Ctx) error {
		return c.JSON(http.StatusOK, bytego.Map{"hello": "world"})
	})
	router.POST("/signupByEmail", server.signupByEmail)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) bytego.Map {
	return bytego.Map{
		"error": err.Error(),
	}
}
