package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	appVersion = "dirty"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 9501, "Prometheus port")
	flag.Parse()

	log.Info("Beginning to serve on port :", port)

	http.Handle("/probe", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
