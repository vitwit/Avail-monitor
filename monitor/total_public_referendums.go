package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchReferendumCount(cfg *config.Config) (string, error) {
	rcendpoint := cfg.Endpoint.URLEndpoint + "/pallets/democracy/storage/referendumCount"
	resp, err := http.Get(rcendpoint)
	if err != nil {
		fmt.Println("failed to fetch referendum count", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch referendum count %d\n", resp.StatusCode)
		return "", err
	}

	var response types.ReferendumCount
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	referendum := response.Value
	return referendum, nil
}
