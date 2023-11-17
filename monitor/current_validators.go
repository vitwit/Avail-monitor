package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchCurrentValidators returns the no. of current validators
func FetchCurrentValidators(cfg *config.Config) (int, error) {
	cvEndpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/session/storage/validators"
	resp, err := http.Get(cvEndpoint)
	if err != nil {
		fmt.Println("Failed to fetch current validator value:", err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch status code of current validator: %d\n", resp.StatusCode)
		return 0, err
	}

	var response types.CurrentValidators
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch data for current validators from endpoint:", err)
		return 0, err
	}

	currentValidators := len(response.Value)
	return currentValidators, nil
}
