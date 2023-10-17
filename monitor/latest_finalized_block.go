package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchFinalizedBlock(cfg *config.Config) (string, error) {
	finalizedendpoint := cfg.Endpoint.URLEndpoint + "/blocks/head"
	resp, err := http.Get(finalizedendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.FinalizedBlock
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	finalizedblock := response.Hash
	return finalizedblock, nil
}
