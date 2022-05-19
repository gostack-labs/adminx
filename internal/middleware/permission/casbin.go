package permission

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/middleware/auth"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgx/v4"
	pgxAdapter "github.com/pckhoi/casbin-pgx-adapter"
)

var Enforcer *casbin.SyncedEnforcer

func Casbin() *casbin.SyncedEnforcer {
	cc, _ := pgx.ParseConfig(configs.Get().DB.Source)
	a, _ := pgxAdapter.NewAdapter(cc, pgxAdapter.WithTableName(configs.Get().Casbin.TableName), pgxAdapter.WithDatabase(configs.Get().Casbin.DBName))
	Enforcer, _ = casbin.NewSyncedEnforcer(configs.Get().Casbin.RbacModel, a)
	_ = Enforcer.LoadPolicy()
	Enforcer.StartAutoLoadPolicy(configs.Get().Casbin.IntervalTime)
	return Enforcer
}

func CheckPermMiddleware() bytego.HandlerFunc {
	return func(c *bytego.Ctx) error {
		obj := c.Request.URL.Path

		act := c.Request.Method

		sub, exist := c.Get(auth.AuthorizationPayloadKey)
		if !exist {
			return resp.Fail(http.StatusUnauthorized, code.RBACError).AbortJSON(c)
		}
		if ok, _ := Enforcer.Enforce(sub, obj, act); !ok {
			return resp.Fail(http.StatusUnauthorized, code.RBACError).AbortJSON(c)
		}
		return c.Next()
	}
}
