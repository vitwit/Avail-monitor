package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchTotalRewardsDistributed(cfg *config.Config) (string, error) {
	currentEra, err := FetchCurrentEra(cfg)
	if err != nil {
		fmt.Println("failed to fetch current era value for total rewards distributed:", err)
		return "", err
	}

	trdendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/staking/storage/erasRewardPoints?keys[]=" + currentEra
	res, err := http.Get(trdendpoint)
	if err != nil {
		fmt.Println("failed to fetch total rewards distributed value", err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch current total rewards distributed code %d\n", res.StatusCode)
		return "", err
	}

	var response types.TotalRewardsDistributed
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal total rewards distributed JSON:", err)
		return "", err
	}
	rewardsdist := response.Value.Total
	return rewardsdist, nil
}
