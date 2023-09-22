package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	router2 "ginweb/application/router"
	"ginweb/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type IApp interface {
	parseCmd()    // 解下cmd参数
	parseConfig() // 解析配置
	initConfig()  // 初始化配置
	runServe()    // 启动http服务
	signalListen(ctx context.Context)
	shutdown(ctx context.Context, ser *http.Server)
	Run()
}

func init() {

}

var (
	sig      = make(chan os.Signal, 1) // 接收停止信号
	server   = &http.Server{}          // http server
	confPath = "./conf.yaml"           // yaml config file path
)

var (
	BuildVersion = ""
	BuildTag     = ""
	BuildCommit  = ""
	BuildTime    = ""
	GoVersion    = ""
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		err := config.AppLogger.Sync()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	parseCmd()
	parseConfig()
	initConfig()
	runServe()
	signalListen(ctx)
}

func parseCmd() {
	var version bool
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&confPath, "conf", confPath, "yaml config file path")
	flag.Parse()
	if version {
		fmt.Printf(" build version: %s\n build tag: %s\n build commit: %s\n build time: %s\n go version: %s\n",
			BuildVersion, BuildTag, BuildCommit, BuildTime, GoVersion)
		os.Exit(0)
		return
	}
	log.Println("parse cmd success")
}

func parseConfig() {
	if err := config.ParseConfig(confPath); err != nil {
		fmt.Printf("parse yaml config: %s , error: %s\n", confPath, err.Error())
		os.Exit(1)
		return
	}
	log.Println("parse yaml config success")
}

func initConfig() {
	log.Println("init config success")
}

func runServe() {

	if config.AppConfig.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	e := gin.Default()
	e.Use(router2.Cors())
	router2.Routes(e)

	port := config.AppConfig.Server.Port
	server.Addr = fmt.Sprintf(":%s", port)
	server.Handler = e
	go func() {
		log.Println(fmt.Sprintf("start serve port: %s", port))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("start serve port: %s, error: %s\n", port, err)
			return
		}
	}()
}

func signalListen(ctx context.Context) {
	log.Println("listen os signal...")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	run := true
	for run {
		select {
		case <-sig:
			run = false
			shutdown(ctx, server)
		}
	}
	log.Println("stop serve")
}

func shutdown(ctx context.Context, ser *http.Server) {
	err := ser.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("serves shutdown success")
}
