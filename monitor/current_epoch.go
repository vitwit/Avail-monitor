package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchEpochIndex returns the current epoch
func FetchEpochIndex(cfg *config.Config) (string, error) {
	epochendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/babe/storage/epochIndex"
	resp, err := http.Get(epochendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.EpochIndex
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch current epoch from endpoint:", err)
		return "", err
	}

	value := response.Value
	return value, nil
}
