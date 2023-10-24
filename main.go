package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/exporter"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	collector := exporter.NewAvailCollector(cfg)

	go collector.WatchSlots(cfg)

	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler()) // exported metrics can be seen in /metrics
	err = http.ListenAndServe(fmt.Sprintf("%s", cfg.Prometheus.ListenAddress), nil)
	if err != nil {
		log.Printf("Error while listening on server : %v", err)
	}

}
