package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gostack-labs/adminx/configs"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
)

// Server serves HTTP requests for our adminx service.
type Server struct {
	config configs.AppConfig
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and set up routing.
func NewServer() (*Server, error) {
	server := &Server{
		config: *configs.Cfg,
		store:  db.NewStore(),
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/signupByEmail", server.signupByEmail)

	router.NoRoute(func(ctx *gin.Context) {
		acceptStr := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 URL 和请求方法是否正确。"})
		}
	})

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(configs.Cfg.Server.Addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
