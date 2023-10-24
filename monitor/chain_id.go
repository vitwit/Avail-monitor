package monitor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
)

func FetchChainID(cfg *config.Config) (string, error) {
	endpoint := cfg.RPC_Endpoint.URLEndpoint + "/node/version"
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch data. Status code: %d\n", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", err
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to fetch data for chain id from endpoint:", err)
		return "", err
	}

	chain, found := data["chain"]
	if !found {
		fmt.Println("Chain not found in response")
		return "", err
	}

	return chain, nil
}
