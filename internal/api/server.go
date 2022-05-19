package api

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/middleware/auth"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/repository/redis"
	"github.com/gostack-labs/adminx/pkg/token"
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
	tokenMaker, err := token.NewPasetoMaker(configs.Get().Token.Key)
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

var (
	utrans   *ut.UniversalTranslator
	validate *validator.Validate
)

func (server *Server) setupRouter() {
	router := bytego.New()
	router.Debug(true)
	router.Use(logger.New())
	router.Use(recovery.New())
	//router.SetValidator(validator.New().Struct)
	// if err := v.InitTrans(router, "zh"); err != nil {
	// 	log.Fatal("translator err:", err)
	// }
	validate = validator.New()
	en := en.New()
	zh := zh.New()
	utrans = ut.New(zh, en)
	trans, _ := utrans.GetTranslator("zh")
	_ = zh_trans.RegisterDefaultTranslations(validate, trans)
	_ = validate.RegisterValidation("phone", validPhone)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		count := 2
		name := strings.SplitN(field.Tag.Get("json"), ",", count)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	router.SetValidator(validate.Struct)
	router.Use(i18n())

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
	user.DELETE("/single/:username", server.deleteUser)
	user.DELETE("/batch", server.batchDeleteUser)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

var validPhone validator.Func = func(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		result, err := regexp.MatchString(`^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`, phone)
		if err != nil {
			return false
		}
		return result
	}
	return true
}

func i18n() bytego.HandlerFunc {
	return func(c *bytego.Ctx) error {
		locale := c.Query("locale")
		var t ut.Translator
		if len(locale) > 0 {
			var found bool
			switch locale {
			case "zh-CN":
				if t, found = utrans.GetTranslator("zh"); found {
					goto END
				}
			case "en-US":
				if t, found = utrans.GetTranslator("en"); found {
					goto END
				}
			default:
				if t, found = utrans.GetTranslator(locale); found {
					goto END
				}
			}

		}
		t, _ = utrans.FindTranslator(acceptedLanguages(c.Request)...)
	END:
		c.Set("transKey", t)
		return c.Next()
	}
}

func acceptedLanguages(r *http.Request) (languages []string) {
	accepted := r.Header.Get("Accept-Language")
	if accepted == "" {
		return
	}
	options := strings.Split(accepted, ",")
	l := len(options)
	languages = make([]string, l)

	for i := 0; i < l; i++ {
		locale := strings.SplitN(options[i], ";", 2)
		s := strings.Trim(locale[0], " ")
		switch s {
		case "zh-CN":
			languages[i] = "zh"
		case "en-US":
			languages[i] = "en"
		default:
			languages[i] = "zh"
		}
	}
	return
}
