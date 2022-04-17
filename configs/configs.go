package configs

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gostack-labs/adminx/pkg/env"
	"github.com/gostack-labs/adminx/pkg/file"
	"github.com/spf13/viper"
)

var Config config

type config struct {
	App        app
	Server     server
	DB         db
	Redis      redis
	Token      token
	Mail       mail
	VerifyCode verifycode
}

type app struct {
	Name    string
	Version string
	Mode    string
	SSL     bool
	CSRF    bool
	Debug   bool
	Welcome string
}

type server struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type db struct {
	Driver string
	Source string
}

type redis struct {
	Addr         string
	Pass         string
	Db           int
	MaxRetries   int
	PoolSize     int
	MinIdleConns int
}

type token struct {
	Key                  string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type mail struct {
	Host string
	Port int
	User string
	Pass string
	To   string
}

type verifycode struct {
	KeyPrefix  string
	ExpireTime int64
}

var once sync.Once

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

func LoadConfig() {
	once.Do(func() {
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

		err := viper.ReadConfig(r)
		if err != nil {
			log.Fatal("read config err:", err)
		}

		err = viper.Unmarshal(&Config)
		if err != nil {
			log.Fatal("viper.Unmarshal err:", err)
		}

		viper.SetConfigName(env.Active().Value() + "_configs")
		viper.AddConfigPath("./configs")

		configFile := "./configs/" + env.Active().Value() + "_configs.yaml"
		_, ok := file.IsExists(configFile)
		if !ok {
			err = os.MkdirAll(filepath.Dir(configFile), os.ModePerm)
			if err != nil {
				log.Fatal("mkdir err:", err)
			}

			f, fErr := os.Create(configFile)
			if fErr != nil {
				log.Fatal("create file err:", fErr)
			}
			defer f.Close()

			err = viper.WriteConfig()
			if err != nil {
				log.Fatal("viper.WriteConfig err:", err)
			}
		}

		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			err = viper.Unmarshal(&Config)
			if err != nil {
				log.Fatal("config change Unmarshal err:", err)
			}
		})
	})
}
