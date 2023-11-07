package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchTotalRewardsClaimed(cfg *config.Config) (string, error) {
	nominationPool, err := FetchNominationPool(cfg)
	if err != nil {
		fmt.Println("failed to fetch nomination pool value for total rewards claimed:", err)
		return "", err
	}

	trcendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/staking/storage/erasRewardPoints?keys[]=" + nominationPool
	res, err := http.Get(trcendpoint)
	if err != nil {
		fmt.Println("failed to fetch total rewards claimed value", err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch current total rewards claimed code %d\n", res.StatusCode)
		return "", err
	}

	var response types.RewardPoolClaimed
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal total rewards claimed JSON:", err)
		return "", err
	}
	rewardclaim := response.Value.TotalRewardsClaimed
	return rewardclaim, nil
}
