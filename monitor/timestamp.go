package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchTimeStamp(cfg *config.Config) (string, error) {
	tsendpoint := cfg.RPC_Endpoint.URLEndpoint + "/blocks/head"
	resp, err := http.Get(tsendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch timestamp", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.TimeStamp
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch timestamp from endpoint:", err)
		return "", err
	}
	epochtime := response.Extrinsics[0].Args.Now
	return epochtime, nil
}
