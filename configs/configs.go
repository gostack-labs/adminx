package configs

import (
	"log"
	"time"

	"github.com/gostack-labs/adminx/pkg/config"
	"github.com/gostack-labs/adminx/pkg/env"
)

var Cfg = new(AppConfig)

type AppConfig struct {
	App
	Server
	DB
	Redis
	Token
	Mail
}

type App struct {
	Name    string
	Version string
	Mode    string
	SSL     bool
	CSRF    bool
	Debug   bool
	Welcome string
}

type Server struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DB struct {
	Driver string
	Source string
}

type Redis struct {
	Addr         string
	Pass         string
	Db           int
	MaxRetries   int
	PoolSize     int
	MinIdleConns int
}

type Token struct {
	Key                  string
	AccessTokenDuration  string
	RefreshTokenDuration string
}

type Mail struct {
	Host string
	Port int
	User string
	Pass string
	To   string
}

func Boot() {}

func init() {
	configFileType := "yaml"
	configFileName := env.Active().Value() + "." + configFileType

	c := config.New("./configs", config.WithFileType(configFileType))
	err := c.Load(configFileName, &Cfg)
	if err != nil {
		log.Fatal("connot load config:", err)
	}
}
