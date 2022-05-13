package main

import (
	"net/http"
	"strconv"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/service"
)

func main() {
	srv, err := service.New("config.json")
	if err != nil {
		logger.Warning.Fatal("Service initialization failed ", err)
	}

	logger.Info.Println("Connecting to port ", srv.Cfg.Port)

	err = http.ListenAndServe(":"+strconv.Itoa(srv.Cfg.Port), routes(srv.Handlers))
	logger.Warning.Fatal(err)
}
