package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchEpochEndTime(cfg *config.Config) (string, error) {
	epochendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/babe/storage/epochStart"
	resp, err := http.Get(epochendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch end time", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch end time code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.EpochEndTime
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch current epoch end time from endpoint:", err)
		return "", err
	}

	endTime := response.Value[1]
	return endTime, nil
}
