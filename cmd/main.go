package main

import (
	"fmt"
	"net/http"
	"os"
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

	writePID()

	err = http.ListenAndServe(":"+strconv.Itoa(srv.Cfg.Port), routes(srv.Handlers))
	logger.Warning.Fatal(err)
}

func writePID() {
	pid := os.Getpid()

	f, err := os.Create("filter_service.pid")

	if err != nil {
		logger.Warning.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("%d", pid))

	if err2 != nil {
		logger.Warning.Fatal(err2)
	}
}
