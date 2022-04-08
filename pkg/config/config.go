package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf *Config

type Config struct {
	env        string
	configDir  string
	configType string // file type, eg: yaml, json, toml, default is yaml
	val        map[string]*viper.Viper
	mu         sync.Mutex
}

// New create a config instance.
func New(cfgDir string, opts ...Option) *Config {
	// must set config dir
	if cfgDir == "" {
		panic("config dir is not set")
	}
	c := Config{
		configDir:  cfgDir,
		configType: "yaml",
		val:        make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(&c)
	}

	conf = &c

	return &c
}

// Load alias for config func.
func Load(filename string, val interface{}) error { return conf.Load(filename, val) }

// Load scan data to struct.
func (c *Config) Load(filename string, val interface{}) error {
	v, err := c.LoadWithtype(filename, c.configType)
	if err != nil {
		return err
	}
	err = v.Unmarshal(&val)
	if err != nil {
		return err
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&val); err != nil {
			panic(err)
		}
		log.Printf("config file changed: %s", in.Name)
	})
	return nil
}

// LoadWithtype load conf by file type.
func LoadWithtype(filename string, cfgType string) (*viper.Viper, error) {
	return conf.LoadWithtype(filename, cfgType)
}

// LoadWithType load conf by file type.
func (c *Config) LoadWithtype(filename string, cfgType string) (*viper.Viper, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.val[filename]
	if ok {
		return v, nil
	}

	v, err := c.load(filename, cfgType)
	if err != nil {
		return nil, err
	}
	c.val[filename] = v
	return v, nil
}

// Load load file.
func (c *Config) load(filename string, cfgType string) (*viper.Viper, error) {
	// application parameters take precedence over environment variables
	env := GetEnvString("APP_ENV", "")
	path := filepath.Join(c.configDir, env)
	if c.env != "" {
		path = filepath.Join(c.configDir, c.env)
	}

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType(c.configType)
	if cfgType != "" {
		v.SetConfigType(cfgType)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	v.WatchConfig()
	return v, nil
}

// GetEnvString get value from env.
func GetEnvString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
