package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchCurrentValidators(cfg *config.Config) ([]string, error) {
	cvEndpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/session/storage/validators"
	resp, err := http.Get(cvEndpoint)
	if err != nil {
		fmt.Println("Failed to fetch current validator value:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch status code of current validator: %d\n", resp.StatusCode)
		return nil, err
	}

	var response types.CurrentValidators
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch data for current validators from endpoint:", err)
		return nil, err
	}

	currentValidators := response.Value
	return currentValidators, nil
}
