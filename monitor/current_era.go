package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchCurrentEra returns the current era of the network
func FetchCurrentEra(cfg *config.Config) (string, error) {
	eraendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/staking/storage/currentEra"
	resp, err := http.Get(eraendpoint)
	if err != nil {
		fmt.Println("failed to fetch current era value", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch current era value %d\n", resp.StatusCode)
		return "", err
	}

	var response types.CurrentEra
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch current epoch from endpoint:", err)
		return "", err
	}

	value := response.Value
	return value, nil
}
