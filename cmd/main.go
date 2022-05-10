package main

import (
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/service"
	"log"
	"net/http"
	"strconv"
)

const portNumber = ":8080"

func main() {

	srv, err := service.New("config.json")
	if err != nil {
		log.Fatal("Service initialization failed ", err)
	}

	logger.Info.Println("Connecting to port ", srv.Cfg.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(srv.Cfg.Port), routes(srv.Handlers))
	log.Fatal(err)
}
