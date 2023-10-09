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
	apiEndpoint = "http://64.227.177.52:8080"
	metricsPort = 9000
)

var (
	nodeVersion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "node_version",
			Help: "Node Version Information",
		},
		[]string{"version"},
	)
	chainName = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "chain",
			Help: "Name of the chain",
		},
		[]string{"chain"},
	)
)

func fetchDataAndSetMetric() {
	endpoint := apiEndpoint + "/node/version"
	fmt.Printf("apiEndpoint: %v\n", endpoint)
	resp, err := http.Get(endpoint)
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

	nodeVersion.WithLabelValues(version).Set(1.0) // Use a constant value (1) for the metric
	chainName.WithLabelValues(chain).Set(1)
	fmt.Printf("chain: %s\n", chain)
	fmt.Printf("Node Version: %s\n", version)
}

func main() {
	version, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	prometheus.MustRegister(nodeVersion)
	prometheus.MustRegister(chainName)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
	}()

	for {
		fetchDataAndSetMetric()
		time.Sleep(25 * time.Second)
	}

	fmt.Println(string(version))
}

//----------------------------------
