package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchFinalizedBlock(cfg *config.Config) {
// 	finalizedendpoint := cfg.URLEndpoint + "/blocks/head"
// 	resp, err := http.Get(finalizedendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.FinalizedBlock
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	finalizedblock := response.Hash
// 	h, _ := strconv.ParseFloat(finalizedblock, 64)
// 	fmt.Printf("finalized block ***********%v\n", h)
// 	finalizedBlock.Set(h)

// }
