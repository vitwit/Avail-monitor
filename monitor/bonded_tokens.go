package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchBondedToken(cfg *config.Config) (string, error) {
	currentEra, err := FetchCurrentEra(cfg)
	if err != nil {
		fmt.Println("failed to fetch current era value for bonded token:", err)
		return "", err
	}

	btendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/staking/storage/erasTotalStake?keys[]=" + currentEra
	res, err := http.Get(btendpoint)
	if err != nil {
		fmt.Println("failed to fetch bonded token value", err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch current bonded token code %d\n", res.StatusCode)
		return "", err
	}

	var response types.BondedTokens
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal bonded token JSON:", err)
		return "", err
	}
	bonded := response.Value
	return bonded, nil
}
