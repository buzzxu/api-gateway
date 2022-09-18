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
	conf.LoadDefaultConf()
	logger.InitLogger()
	ironman.Server(router.New())
	flag.Usage = usage
	flag.Usage()
}

func usage() {
	fmt.Println("Api Gateway v0.01")
}
