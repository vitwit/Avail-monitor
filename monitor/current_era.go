package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchCurrentEra(cfg *config.Config) {
// 	eraendpoint := cfg.URLEndpoint + "/pallets/staking/storage/currentEra"
// 	fmt.Printf("currentSlot: %v\n", eraendpoint)
// 	resp, err := http.Get(eraendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch current era value", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current era value %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CurrentEra
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}
// 	value := response.Value
// 	e, _ := strconv.ParseFloat(value, 64)
// 	currentEra.Set(e)
// }
