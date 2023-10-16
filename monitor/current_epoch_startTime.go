package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchEpochStartTime(cfg *config.Config) {
// 	startendpoint := cfg.URLEndpoint + "/pallets/babe/storage/epochStart"
// 	fmt.Printf("epoch start time: %v\n", startendpoint)
// 	resp, err := http.Get(startendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch start time", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch start time code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochStartTime
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	startTime := response.Value[0]
// 	fmt.Println(startTime)
// 	st, _ := strconv.ParseFloat(startTime, 64)
// 	epochstartTime.Set(st)

// }
