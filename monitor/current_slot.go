package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchCurrentSlot returns the current slot of the network
func FetchCurrentSlot(cfg *config.Config) (string, error) {
	finalendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/babe/storage/currentSlot"
	resp, err := http.Get(finalendpoint)
	if err != nil {
		fmt.Println("failed to fetch finalied block", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch finalzed code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.CurrentSlot
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch current slot from endpoint:", err)
		return "", err
	}

	value := response.Value
	return value, nil
}
