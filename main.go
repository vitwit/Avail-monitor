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
	apiEndpoint  = "http://64.227.177.52:8080"
	metricsPort  = 9000
	timeInterval = 25 * time.Second
)

type CurrentSlot struct {
	Value string `json:"value"`
}

type EpochIndex struct {
	Value string `json:"value"`
}

type TimeStamp struct {
	Extrinsics []struct {
		Args struct {
			Now string `json:"now"`
		} `json:"args"`
	} `json:"extrinsics"`
}

type BestBlock struct {
	Number string `json:"number"`
}

type FinalizedBlock struct {
	Hash string `json:"hash"`
}

type EpochStartTime struct {
	Value []string `json:"value"`
}

type EpochEndTime struct {
	Value []string `json:"value"`
}

type TotalTokensIssued struct {
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
	currentSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_slot_value",
		Help: "current slot value of avail",
	},
	)
	epochIndex = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_epoch_number",
		Help: "current epoch number of avail",
	},
	)
	timeStamp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "timestamp_of_latest_block",
		Help: "timestamp of latest block",
	},
	// []string{"now"},
	)
	bestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "latest_best_block",
		Help: "latest best block",
	},
	)
	finalizedBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "finalized_block",
		Help: "finalized block of the network",
	},
	)
	epochstartTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "epoch_start_time",
		Help: "epoch start time of network",
	},
	)
	epochendTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "epoch_end_time",
		Help: "epoch end time of network",
	},
	)
	totaltokensIssued = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_tokens_issued",
		Help: "total tokens issued on network",
	})
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

	nodeVersion.WithLabelValues(version).Set(1)

	// nodeVersion.WithLabelValues(version).Set(n) // Use a constant value (1) for the metric
	chainName.WithLabelValues(chain).Set(1)
	fmt.Printf("chain: %s\n", chain)
	fmt.Printf("Node Version: %s\n", version)
}

func fetchCurrentSlot() {
	finalendpoint := apiEndpoint + "/pallets/babe/storage/currentSlot"
	fmt.Printf("currentSlot: %v\n", finalendpoint)
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

	var response CurrentSlot
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	value := response.Value

	v, _ := strconv.ParseFloat(value, 64)
	fmt.Println("value here....", v)
	currentSlot.Set(v)
	fmt.Printf("Finalized Block Value: %s\n", value)

}

func fetchEpochIndex() {
	epochendpoint := apiEndpoint + "/pallets/babe/storage/epochIndex"
	fmt.Printf("epochindex enddpoint: %v\n", epochendpoint)
	resp, err := http.Get(epochendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return
	}

	var response EpochIndex
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	value := response.Value
	e, _ := strconv.ParseFloat(value, 64)
	epochIndex.Set(e)
	fmt.Printf("epoch index value: %s\n", value)

}

func fetchTimeStamp() {
	tsendpoint := apiEndpoint + "/blocks/head"
	resp, err := http.Get(tsendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch timestamp", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return
	}

	var response TimeStamp
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	// fmt.Printf("response: %v\n", response)

	// fmt.Printf("response.Extrinsics: %v\n", response.Extrinsics)

	// fmt.Printf("response.Extrinsics[0].Args.Now: %v\n", response.Extrinsics[0].Args.Now)

	epochtime := response.Extrinsics[0].Args.Now
	fmt.Printf("epochtime: %v\n", epochtime)

	// ts, err := strconv.ParseInt(epochtime, 10, 64)
	// if err != nil {
	// 	fmt.Println("Error parsing epoch value:", err)
	// 	return
	// }

	// t := time.Unix(ts/1000, 0) // Convert milliseconds to seconds
	// loc, err := time.LoadLocation("GMT")
	// if err != nil {
	// 	fmt.Println("Error loading timezone:", err)
	// 	return
	// }
	// t = t.In(loc)
	// formattedTime := t.Format("Monday, January 02, 2006 3:04:05 PM MST")
	ts, _ := strconv.ParseFloat(epochtime, 64)
	// Export the converted timestamp to Prometheus
	timeStamp.Set(ts) // Export as seconds

	fmt.Printf("Fetched timestamp: %s\n", epochtime)
}

func fetchBestBlock() {
	blockendpoint := apiEndpoint + "/blocks/head"
	fmt.Printf("epochindex enddpoint: %v\n", blockendpoint)
	resp, err := http.Get(blockendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return
	}

	var response BestBlock
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	block := response.Number
	b, _ := strconv.ParseFloat(block, 64)
	bestBlock.Set(b)
}

func fetchFinalizedBlock() {
	finalizedendpoint := apiEndpoint + "/blocks/head"
	resp, err := http.Get(finalizedendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return
	}

	var response FinalizedBlock
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	finalizedblock := response.Hash
	h, _ := strconv.ParseFloat(finalizedblock, 64)
	finalizedBlock.Set(h)

}

func fetchEpochStartTime() {
	startendpoint := apiEndpoint + "/pallets/babe/storage/epochStart"
	fmt.Printf("epoch start time: %v\n", startendpoint)
	resp, err := http.Get(startendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch start time", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch start time code %d\n", resp.StatusCode)
		return
	}

	var response EpochStartTime
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	startTime := response.Value[0]
	fmt.Println(startTime)
	st, _ := strconv.ParseFloat(startTime, 64)
	epochstartTime.Set(st)

}

func fetchEpochEndTime() {
	epochendpoint := apiEndpoint + "/pallets/babe/storage/epochStart"
	fmt.Printf("epoch end time: %v\n", epochendpoint)
	resp, err := http.Get(epochendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch end time", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch end time code %d\n", resp.StatusCode)
		return
	}

	var response EpochEndTime
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	endTime := response.Value[1]
	fmt.Println(endTime)
	et, _ := strconv.ParseFloat(endTime, 64)
	epochendTime.Set(et)

}

func fetchTotalTokensIssued() {}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	// version, err := os.ReadFile("config.toml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	prometheus.MustRegister(nodeVersion)
	prometheus.MustRegister(chainName)
	prometheus.MustRegister(currentSlot)
	prometheus.MustRegister(epochIndex)
	prometheus.MustRegister(timeStamp)
	prometheus.MustRegister(bestBlock)
	prometheus.MustRegister(finalizedBlock)
	prometheus.MustRegister(epochstartTime)
	prometheus.MustRegister(epochendTime)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
	}()

	for {
		<-ticker.C
		fetchDataAndSetMetric()
		fetchCurrentSlot()
		fetchEpochIndex()
		fetchTimeStamp()
		fetchBestBlock()
		fetchFinalizedBlock()
		fetchEpochStartTime()
		fetchEpochEndTime()
		time.Sleep(timeInterval)
	}

	// fmt.Println(string(version))
}
