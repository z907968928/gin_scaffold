package router

import (
	"context"
	"fmt"
	"github.com/e421083458/gin_scaffold/core"
	"github.com/e421083458/gin_scaffold/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(core.ConfBase.Base.DebugMode)
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           core.ConfBase.HTTP.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(core.ConfBase.HTTP.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(core.ConfBase.HTTP.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(core.ConfBase.HTTP.MaxHeaderBytes),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", core.ConfBase.HTTP.Addr)
		pid := fmt.Sprint(os.Getpid())
		log.Printf("pid: %s", pid)
		_ = utils.WriteFile(core.ConfBase.HTTP.PidPath, []byte(pid), 0666)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", core.ConfBase.HTTP.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
