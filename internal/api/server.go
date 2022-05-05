package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/middleware/auth"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/repository/redis"
	"github.com/gostack-labs/adminx/pkg/token"
	v "github.com/gostack-labs/adminx/pkg/validate"
	"github.com/gostack-labs/bytego"
	"github.com/gostack-labs/bytego/middleware/logger"
	"github.com/gostack-labs/bytego/middleware/recovery"
)

// Server serves HTTP requests for our adminx service.
type Server struct {
	store      db.Store
	cache      redis.Store
	tokenMaker token.Maker
	router     *bytego.App
}

// NewServer create a new HTTP server and set up routing.
func NewServer(store db.Store, cache redis.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(configs.Config.Token.Key)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		cache:      cache,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := bytego.New()
	router.Debug(true)
	router.Use(logger.New())
	router.Use(recovery.New())
	//router.Validator(validator.New().Struct)
	if err := v.InitTrans(router, "zh"); err != nil {
		log.Fatal("translator err:", err)
	}
	router.GET("/", func(c *bytego.Ctx) error {
		return c.JSON(http.StatusOK, bytego.Map{"hello": "world"})
	})
	router.POST("/signup", server.signup)
	router.POST("/signup/sendUsingEmail", server.sendUsingEmail)

	router.POST("/signin", server.logginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	sys := router.Group("/sys", auth.AuthMiddleware(server.tokenMaker))
	menu := sys.Group("/menu")
	menu.GET("/tree", server.menuTree)
	menu.POST("", server.createMenu)
	menu.PUT("/:id", server.updateMenu)
	menu.DELETE("/single/:id", server.deleteMenu)
	menu.DELETE("/batch", server.batchDeleteMenu)
	menu.GET("/button/:id", server.menuButton)
	menu.POST("/api/:id", server.mentBindApi)
	menu.GET("/api/:menu", server.MenuApis)
	menu.GET("/api-list/:menu", server.MenuApiList)

	apiGroup := sys.Group("/api-group")
	apiGroup.GET("", server.listApiGroup)
	apiGroup.POST("", server.createApiGroup)
	apiGroup.PUT("/:id", server.updateApiGroup)
	apiGroup.DELETE("/single/:id", server.deleteApiGroup)
	apiGroup.DELETE("/batch", server.batchDeleteApiGroup)

	api := sys.Group("/api")
	api.GET("", server.listApi)
	api.POST("", server.createApi)
	api.PUT("/:id", server.updateApi)
	api.DELETE("/single/:id", server.deleteApi)
	api.DELETE("/batch", server.batchDeleteApi)

	role := sys.Group("/role")
	role.GET("", server.listRole)
	role.POST("", server.createRole)
	role.PUT("/:id", server.updateRole)
	role.DELETE("/single/:id", server.deleteRole)
	role.DELETE("/batch", server.batchDeleteRole)
	role.POST("/permission/:id", server.updateRolePermission)
	role.POST("/api/:id", server.roleApiPermission)
	role.GET("/api/:id", server.getRoleApi)
	role.GET("permission/:id", server.getRolePermission)

	user := sys.Group("/user")
	user.GET("", server.listUser)
	user.GET("/info", server.userInfo)
	user.GET("/info/:username", server.userInfoByID)
	user.POST("", server.createUser)
	user.PUT("/:username", server.updateUser)
	user.DELETE("/:username", server.deleteUser)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) bytego.Map {
	if errs, ok := err.(validator.ValidationErrors); ok {
		rsp := make(bytego.Map)
		for field, terr := range errs.Translate(v.Trans) {
			rsp[field[strings.Index(field, ".")+1:]] = terr
		}
		return bytego.Map{
			"error": rsp,
		}
	}
	return bytego.Map{
		"error": err.Error(),
	}
}
