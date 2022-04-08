package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Environment
	local  Environment = &environment{value: "local"}
	test   Environment = &environment{value: "test"}
	stage  Environment = &environment{value: "stage"}
	pro    Environment = &environment{value: "pro"}
)

// Environment 环境配置
type Environment interface {
	Value() string
	IsLocal() bool
	IsTest() bool
	IsStage() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsLocal() bool {
	return e.value == "local"
}

func (e *environment) IsTest() bool {
	return e.value == "test"
}

func (e *environment) IsStage() bool {
	return e.value == "stage"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n local:开发环境\n test:测试环境\n stage:预上线环境\n pro:正是环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "local":
		active = local
	case "test":
		active = test
	case "stage":
		active = stage
	case "pro":
		active = pro
	default:
		active = stage
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'stage' will be used.")
	}
}

// Active 当前配置的 env
func Active() Environment {
	return active
}
