package configs

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gostack-labs/adminx/pkg/env"
	"github.com/gostack-labs/adminx/pkg/file"
	"github.com/spf13/viper"
)

type Config struct {
	App    App
	Server Server
	DB     DB
	Redis  Redis
	Token  Token
	Mail   Mail
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

var (
	//go:embed local_configs.yaml
	localConfigs []byte

	//go:embed test_configs.yaml
	testConfigs []byte

	//go:embed stage_configs.yaml
	stageConfigs []byte

	//go:embed pro_configs.yaml
	proConfigs []byte
)

func LoadConfig() (config Config, err error) {
	var r io.Reader

	switch env.Active().Value() {
	case "local":
		r = bytes.NewReader(localConfigs)
	case "test":
		r = bytes.NewReader(testConfigs)
	case "stage":
		r = bytes.NewReader(stageConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(localConfigs)
	}

	viper.SetConfigType("yaml")

	err = viper.ReadConfig(r)
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.AddConfigPath("./configs")

	configFile := "./configs/" + env.Active().Value() + "_configs.yaml"
	_, ok := file.IsExists(configFile)
	if !ok {
		err = os.MkdirAll(filepath.Dir(configFile), os.ModePerm)
		if err != nil {
			return
		}

		f, fErr := os.Create(configFile)
		if fErr != nil {
			err = fErr
			return
		}
		defer f.Close()

		err = viper.WriteConfig()
		if err != nil {
			return
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&config)
		if err != nil {
			return
		}
	})
	return
}
