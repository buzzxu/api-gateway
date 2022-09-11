package main

import (
	"flag"
	"fmt"
	"github.com/buzzxu/ironman"
	"github.com/buzzxu/ironman/conf"
	"github.com/labstack/echo/v4"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(conf.ServerConf.MaxProc)
	flag.Parse()
	e := echo.New()

	ironman.Server(e)
	flag.Usage = usage
	flag.Usage()
}

func usage() {
	fmt.Println("Api Gateway v0.01")
}
