package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchTotalTokensIssued returns the total no of tokens
// issued on the avail network
func FetchTotalTokensIssued(cfg *config.Config) (string, error) {
	tokenendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/balances/storage/totalIssuance"
	resp, err := http.Get(tokenendpoint)
	if err != nil {
		fmt.Println("failed to fetch total token issuance", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch total token issuance code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.TotalTokensIssued
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch total tokens issued from the network:", err)
		return "", err
	}

	totalTokens := response.Value
	return totalTokens, nil
}
