package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchEpochIndex(cfg *config.Config) {
// 	epochendpoint := cfg.URLEndpoint + "/pallets/babe/storage/epochIndex"
// 	fmt.Printf("epochindex enddpoint: %v\n", epochendpoint)
// 	resp, err := http.Get(epochendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochIndex
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	value := response.Value
// 	e, _ := strconv.ParseFloat(value, 64)
// 	epochIndex.Set(e)
// 	fmt.Printf("epoch index value: %s\n", value)

// }
