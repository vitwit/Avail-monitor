package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchCouncilMember(cfg *config.Config) {
// 	cmendpoint := cfg.URLEndpoint + "/pallets/council/storage/members"
// 	resp, err := http.Get(cmendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch nomination pools", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch number of nomination pools %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CouncilMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	councilmem := response.Value[0]
// 	// fmt.Println(councilmem)
// 	// cm, _ := strconv.ParseFloat(councilmem, 32)
// 	// fmt.Printf("cm prom metric-------- %v\n", cm)
// 	councilMember.WithLabelValues(councilmem).Set(1)

// }
