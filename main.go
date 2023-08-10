package main

import (
	"github.com/e421083458/gin_scaffold/initialize"
	"github.com/e421083458/gin_scaffold/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initialize.InitModule("./conf/dev/", []string{
		"base",
		"mysql_map",
		"redis_map",
		"profile",
	}); err != nil {
		panic("init module err")
	}

	defer initialize.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
