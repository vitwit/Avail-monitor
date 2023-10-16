package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchTotalTokensIssued(cfg *config.Config) {
// 	tokenendpoint := cfg.URLEndpoint + "/pallets/balances/storage/totalIssuance"
// 	resp, err := http.Get(tokenendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch total token issuance", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch total token issuance code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.TotalTokensIssued
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	totalTokens := response.Value
// 	tt, _ := strconv.ParseFloat(totalTokens, 64)
// 	abcd := math.Floor(tt / math.Pow(10, 18))
// 	fmt.Printf("abcd: %v\n", abcd)

// 	// ttI := tt / 1e18 //wrong conversion.. consider later..
// 	//totaltokensIssued.Set(abcd)
// 	totaltokensIssued.WithLabelValues(fmt.Sprintf("%.11e", abcd)).Set(1)

// }
