package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	apiEndpoint = "http://64.227.177.52:8080"
	metricsPort = 9000
)

type FinalizedBlock struct {
	Value string `json:"value"`
}

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
	finalizedBlock = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "current_slot_value",
		Help: "current slot value of avail",
	},
		[]string{"value"},
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

	// n, err := strconv.ParseFloat(version, 64)
	// fmt.Println("errrr......", err, n)

	nodeVersion.WithLabelValues(version).Desc().String()

	// nodeVersion.WithLabelValues(version).Set(n) // Use a constant value (1) for the metric
	chainName.WithLabelValues(chain).Set(1)
	fmt.Printf("chain: %s\n", chain)
	fmt.Printf("Node Version: %s\n", version)
}

func fetchFinalizedBlock() {
	finalendpoint := apiEndpoint + "/pallets/babe/storage/currentSlot"
	fmt.Printf("finalizedBlock: %v\n", finalendpoint)
	resp, err := http.Get(finalendpoint)
	if err != nil {
		fmt.Println("failed to fetch finalied block", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch finalzed code %d\n", resp.StatusCode)
		return
	}

	var response FinalizedBlock
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	value := response.Value

	v, _ := strconv.ParseFloat(value, 64)
	fmt.Println("value here....", v)
	finalizedBlock.WithLabelValues(value).Set(v)
	fmt.Printf("Finalized Block Value: %s\n", value)

}

func main() {
	// version, err := os.ReadFile("config.toml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	prometheus.MustRegister(nodeVersion)
	prometheus.MustRegister(chainName)
	prometheus.MustRegister(finalizedBlock)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
	}()

	for {
		fetchDataAndSetMetric()
		fetchFinalizedBlock()
		time.Sleep(25 * time.Second)
	}

	// fmt.Println(string(version))
}
