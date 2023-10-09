//------ two metrics below.........

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	apiEndpoint = "http://64.227.177.52:8080/node/version"
	metricsPort = 9090 // Port for Prometheus to scrape metrics
)

var (
	nodeVersion = prometheus.NewGauge( //slots.go
		prometheus.GaugeOpts{
			Name: "node_version",
			Help: "node Version",
		},
	)
)

var (
	chainName = prometheus.NewGauge( //slots.go
		prometheus.GaugeOpts{
			Name: "chain",
			Help: "name of the chain",
		},
	)
)

func fetchDataAndSetMetric() {
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch data. Status code: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	version, found := data["clientVersion"]
	if !found {
		fmt.Println("Version not found in response")
		return
	}

	chain, found := data["chain"]
	if !found {
		fmt.Println("Chain not found in response")
		return
	}

	chainName.Set(1) // Set it to 1 or any value you prefer
	fmt.Printf("chain: %s\n", chain)
	fmt.Printf("Node Version: %s\n", version)
}

func main() {

	version, err := os.ReadFile("config.toml")

	if err != nil {
		log.Fatal(err)
	}

	for {
		fetchDataAndSetMetric()
		time.Sleep(25 * time.Second)
	}
	fmt.Println(string(version))
	prometheus.MustRegister(nodeVersion) //slots.go file

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
	}()

}
