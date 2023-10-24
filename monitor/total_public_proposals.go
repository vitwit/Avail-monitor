package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchPublicProposalCount(cfg *config.Config) (string, error) {
	ppcendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/democracy/storage/publicPropCount"
	resp, err := http.Get(ppcendpoint)
	if err != nil {
		fmt.Println("failed to fetch public proposal count", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch public proposal count%d\n", resp.StatusCode)
		return "", err
	}

	var response types.PublicProposalCount
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch total public proposals from endpoint:", err)
		return "", err
	}

	publicpc := response.Value
	return publicpc, nil
}
