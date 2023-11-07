package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchTotalNominatorPoolRewards(cfg *config.Config) (string, error) {
	nominationPool, err := FetchNominationPool(cfg)
	if err != nil {
		fmt.Println("failed to fetch nomination pool value for total rewards claimed:", err)
		return "", err
	}

	tnprendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/staking/storage/erasRewardPoints?keys[]=" + nominationPool
	res, err := http.Get(tnprendpoint)
	if err != nil {
		fmt.Println("failed to fetch total nomination pool rewards", err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch total nomination reward pool %d\n", res.StatusCode)
		return "", err
	}

	var response types.NominatorPoolReward
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal total nomination pool rewards JSON:", err)
		return "", err
	}

	nominatorpoolrew := response.Value.LastRecordedRewardCounter
	fmt.Println("******************nomina*************************************", nominatorpoolrew)
	return nominatorpoolrew, nil
}
