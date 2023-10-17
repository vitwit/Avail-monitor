package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func FetchCouncilMember(cfg *config.Config) (string, error) {
// 	cmendpoint := cfg.Endpoint.URLEndpoint + "/pallets/council/storage/members"
// 	resp, err := http.Get(cmendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch nomination pools", err)
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch number of nomination pools %d\n", resp.StatusCode)
// 		return "", err
// 	}

// 	var response types.CouncilMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return "", err
// 	}

// 	councilmem := response.Value[0]
// 	return councilmem, nil
// 	// fmt.Println(councilmem)
// 	// cm, _ := strconv.ParseFloat(councilmem, 64)
// 	// fmt.Printf("cm prom metric-------- %v\n", cm)
// 	// councilMember.WithLabelValues(councilmem).Set(1)

// }
