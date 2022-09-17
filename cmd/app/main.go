package main

import (
	"api-gateway/internal/router"
	"flag"
	"fmt"
	"github.com/buzzxu/ironman"
	"github.com/buzzxu/ironman/conf"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(conf.ServerConf.MaxProc)
	flag.Parse()
	conf.LoadDefaultConf()
	ironman.Server(router.New())
	flag.Usage = usage
	flag.Usage()
}

func usage() {
	fmt.Println("Api Gateway v0.01")
}
