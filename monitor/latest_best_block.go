package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchBestBlock(cfg *config.Config) (string, error) {
	blockendpoint := cfg.RPC_Endpoint.URLEndpoint + "/blocks/head"
	resp, err := http.Get(blockendpoint)
	if err != nil {
		fmt.Println("failed to fetch epoch index", err)
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
		return "", nil
	}

	var response types.BestBlock
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to fetch latest best block from endpoint:", err)
		return "", err
	}

	block := response.Number
	return block, err
}
