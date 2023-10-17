package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchBountyProposalCount(cfg *config.Config) (string, error) {
	bpcendpoint := cfg.Endpoint.URLEndpoint + "/pallets/bounties/storage/bountyCount"
	resp, err := http.Get(bpcendpoint)
	if err != nil {
		fmt.Println("failed to fetch bounty proposal count value", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch bounty proposal count value %d\n", resp.StatusCode)
		return "", err
	}

	var response types.BountyProposalCount
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	bountypc := response.Value
	return bountypc, nil
}
