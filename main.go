// //------ two metrics below.........

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// const (
// 	apiEndpoint = "http://64.227.177.52:8080/node/version"
// 	metricsPort = 9090 // Port for Prometheus to scrape metrics
// )

// var (
// 	nodeVersion = prometheus.NewGauge( //slots.go
// 		prometheus.GaugeOpts{
// 			Name: "node_version",
// 			Help: "node Version",
// 		},
// 	)
// )

// var (
// 	chainName = prometheus.NewGauge( //slots.go
// 		prometheus.GaugeOpts{
// 			Name: "chain",
// 			Help: "name of the chain",
// 		},
// 	)
// )

// func fetchDataAndSetMetric() {
// 	resp, err := http.Get(apiEndpoint)
// 	if err != nil {
// 		fmt.Println("Failed to fetch data:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Failed to fetch data. Status code: %d\n", resp.StatusCode)
// 		return
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Failed to read response body:", err)
// 		return
// 	}

// 	var data map[string]string
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	version, found := data["clientVersion"]
// 	if !found {
// 		fmt.Println("Version not found in response")
// 		return
// 	}

// 	chain, found := data["chain"]
// 	if !found {
// 		fmt.Println("Version not found in response")
// 		return
// 	}

// 	chainName.Set(float64(1))
// 	fmt.Printf("chain: %s\n", chain)
// 	fmt.Printf("Node Version: %s\n", version)
// }

// func main() {

// 	version, err := os.ReadFile("config.toml")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(version))
// 	prometheus.MustRegister(nodeVersion) //slots.go file

// 	go func() {
// 		http.Handle("/metrics", promhttp.Handler())
// 		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
// 		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
// 	}()

// 	for {
// 		fetchDataAndSetMetric()
// 		time.Sleep(60 * time.Second)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vitwit/avail-monitor/config"
)

func main() {
	config, err := config.ReadConfig(".")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//fmt.Println(config)

	v, err := os.ReadFile(config.URLEndpoint)
	fmt.Println(v)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
