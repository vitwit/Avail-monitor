package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchNominationPool returns the no. of nomination pools
func FetchNominationPool(cfg *config.Config) (string, error) {
	poolendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/nominationPools/storage/counterForBondedPools"
	resp, err := http.Get(poolendpoint)
	if err != nil {
		fmt.Println("failed to fetch nomination pools", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch number of nomination pools %d\n", resp.StatusCode)
		return "", err
	}

	var response types.NominationPool
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch nomination pool from endpoint:", err)
		return "", err
	}

	nominationpool := response.Value
	return nominationpool, nil
}
