package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func FetchCouncilProposalCount(cfg *config.Config) (string, error) {
// 	cpendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/council/storage/proposalCount"
// 	resp, err := http.Get(cpendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch council proposal endpoint", err)
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch council proposal count%d\n", resp.StatusCode)
// 		return "", err
// 	}

// 	var response types.CouncilProposals
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to fetch total council proposals from endpoint:", err)
// 		return "", err
// 	}

// 	councilP := response.Value
// 	return councilP, nil
// }
