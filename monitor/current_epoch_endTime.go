package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchEpochEndTime(cfg *config.Config) {
// 	epochendpoint := cfg.URLEndpoint + "/pallets/babe/storage/epochStart"
// 	fmt.Printf("epoch end time: %v\n", epochendpoint)
// 	resp, err := http.Get(epochendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch end time", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch end time code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochEndTime
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	endTime := response.Value[1]
// 	fmt.Println(endTime)
// 	et, _ := strconv.ParseFloat(endTime, 64)
// 	epochendTime.Set(et)

// }
