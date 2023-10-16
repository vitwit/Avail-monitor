package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchBountyProposalCount(cfg *config.Config) {
// 	bpcendpoint := cfg.URLEndpoint + "/pallets/bounties/storage/bountyCount"
// 	resp, err := http.Get(bpcendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch bounty proposal count value", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch bounty proposal count value %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.BountyProposalCount
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	bountypc := response.Value
// 	bpc, _ := strconv.ParseFloat(bountypc, 64)
// 	bountyProposalCount.Set(bpc)
// }
