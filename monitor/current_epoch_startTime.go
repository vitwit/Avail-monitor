package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchEPochStartTime returns the block height on
// which current epoch started
func FetchEpochStartTime(cfg *config.Config) (string, error) {
	startendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/babe/storage/epochStart"
	resp, err := http.Get(startendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch start time", err)
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch start time code %d\n", resp.StatusCode)
		return "", nil
	}

	var response types.EpochStartTime
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch current epoch start time from endpoint:", err)
		return "", nil
	}

	startTime := response.Value[0]
	return startTime, nil
}
