package api

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/adminx/pkg/env"
	"github.com/gostack-labs/bytego"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type viewResponse struct {
	MemTotal       string  // 内存总量
	MemUsed        string  // 内存使用量
	MemUsedPercent float64 // 内存使用百分比

	DiskTotal       string  // 存储总量
	DiskUsed        string  // 存储使用量
	DiskUsedPercent float64 // 存储使用百分比

	HostOS   string // 主机系统
	HostName string // 主机名

	CpuName        string  // cpu 名字
	CpuCores       int32   // cpu 核心
	CpuUsedPercent float64 // cpu使用百分比

	GoPath      string // GOPATH
	GoVersion   string // go 版本
	Goroutine   int    // go 协程
	ProjectPath string // 项目路径
	Env         string // 环境变量
	Host        string // 主机
	GoOS        string // GOOS
	GoArch      string // GOARCH

	ProjectVersion    string // 项目版本
	PostgresqlVersion string // pgsql 版本
	RedisVersion      string // redis 版本
}

//@title 仪表盘接口
//@api get /dashboard
//@group basic
//@response 200 resp.resultOK{businesscode=10000,message="获取成功",data=viewResponse}
func (server *Server) dashboard(c *bytego.Ctx) error {
	pgVer, err := server.store.SelectVersion(c.Context())
	if err != nil {
		resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	redisVer := server.cache.Version()

	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(time.Second, false)

	memUsedPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", memInfo.UsedPercent), 64)
	diskUsedPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", diskInfo.UsedPercent), 64)
	var dash = viewResponse{
		PostgresqlVersion: pgVer,
		RedisVersion:      redisVer,

		MemTotal:       fmt.Sprintf("%d GB", memInfo.Total/GB),
		MemUsed:        fmt.Sprintf("%d GB", memInfo.Used/GB),
		MemUsedPercent: memUsedPercent,

		DiskTotal:       fmt.Sprintf("%d GB", diskInfo.Total/GB),
		DiskUsed:        fmt.Sprintf("%d GB", diskInfo.Used/GB),
		DiskUsedPercent: diskUsedPercent,

		HostOS:   fmt.Sprintf("%s(%s) %s", hostInfo.Platform, hostInfo.PlatformFamily, hostInfo.PlatformVersion),
		HostName: hostInfo.Hostname,
	}

	if len(cpuInfo) > 0 {
		dash.CpuName = cpuInfo[0].ModelName
		dash.CpuCores = cpuInfo[0].Cores
	}

	if len(cpuPercent) > 0 {
		dash.CpuUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuPercent[0]), 64)
	}

	dash.GoPath = runtime.GOROOT()
	dash.GoVersion = runtime.Version()
	dash.Goroutine = runtime.NumGoroutine()
	dir, _ := os.Getwd()
	dash.ProjectPath = strings.Replace(dir, "\\", "/", -1)
	dash.Host = c.Request.Host
	dash.Env = env.Active().Value()
	dash.GoOS = runtime.GOOS
	dash.GoArch = runtime.GOARCH
	dash.ProjectVersion = release
	return resp.GetOK(dash).JSON(c)
}
