package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchBestBlock(cfg *config.Config) {
// 	blockendpoint := cfg.URLEndpoint + "/blocks/head"
// 	fmt.Printf("epochindex enddpoint: %v\n", blockendpoint)
// 	resp, err := http.Get(blockendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.BestBlock
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	block := response.Number
// 	b, _ := strconv.ParseFloat(block, 64)
// 	bestBlock.Set(b)
// }
