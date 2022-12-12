package main

import (
	"api-gateway/internal/router"
	"flag"
	"fmt"
	"github.com/buzzxu/ironman"
	"github.com/buzzxu/ironman/conf"
	"github.com/buzzxu/ironman/logger"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(conf.ServerConf.MaxProc)
	flag.Parse()
	// 关闭redis
	defer ironman.Redis.Close()
	logger.InitLogger()
	ironman.Server(router.New())
	flag.Usage = usage
	flag.Usage()
}

func usage() {
	fmt.Println("Api Gateway v0.01")
}
