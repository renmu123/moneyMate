package main

import (
	"fmt"
	"money/config"
	"money/models"
	"money/routers"
	"net/http"
)

func main() {
	fmt.Println(models.DB)
	fmt.Println(config.RunMode)
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HttpPort),
		Handler: router,
		//ReadTimeout:    cg.Server.ReadTimeOut,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
