package permission

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgx/v4"
	pgxAdapter "github.com/pckhoi/casbin-pgx-adapter"
)

var Enforcer *casbin.SyncedEnforcer

func Casbin() *casbin.SyncedEnforcer {
	cc, _ := pgx.ParseConfig(configs.Config.DB.Source)
	a, _ := pgxAdapter.NewAdapter(cc, pgxAdapter.WithTableName(configs.Config.Casbin.TableName), pgxAdapter.WithDatabase(configs.Config.Casbin.DBName))
	Enforcer, _ = casbin.NewSyncedEnforcer(configs.Config.Casbin.RbacModel, a)
	_ = Enforcer.LoadPolicy()
	Enforcer.StartAutoLoadPolicy(configs.Config.Casbin.IntervalTime)
	return Enforcer
}

func CheckPermMiddleware() bytego.HandlerFunc {
	return func(c *bytego.Ctx) error {
		obj := c.Request.URL.Path

		act := c.Request.Method

		sub, exist := c.Get("username")
		if !exist {
			c.AbortWithStatus(http.StatusNoContent)
		}
		if ok, _ := Enforcer.Enforce(sub, obj, act); !ok {
			c.Abort()
		}
		return c.Next()
	}
}
