package bootstrap

import (
	"github.com/gostack-labs/adminx/configs"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
)

func Initialize() {
	// iniialize config
	configs.Boot()
	db.NewStore()
}
