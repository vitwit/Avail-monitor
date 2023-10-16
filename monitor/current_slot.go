package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchCurrentSlot(cfg *config.Config) {
// 	finalendpoint := cfg.URLEndpoint + "/pallets/babe/storage/currentSlot"
// 	fmt.Printf("currentSlot: %v\n", finalendpoint)
// 	resp, err := http.Get(finalendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch finalied block", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch finalzed code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CurrentSlot
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	value := response.Value

// 	v, _ := strconv.ParseFloat(value, 64)
// 	fmt.Println("value here....", v)
// 	currentSlot.Set(v)
// 	fmt.Printf("current slot Value: %s\n", value)

// }
