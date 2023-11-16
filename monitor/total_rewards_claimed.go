package monitor

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

// FetchTotalRewardsClaimed returns the summation of all reward pool
// by querying each reward pool on the avail network
func FetchTotalRewardsClaimed(cfg *config.Config) (float64, error) {
	nominationPool, err := FetchNominationPool(cfg)
	if err != nil {
		fmt.Println("failed to fetch nomination pool value for total rewards claimed:", err)
		return 0, err
	}

	nomination, err := strconv.ParseInt(nominationPool, 10, 64)
	if err != nil {
		return 0, err
	}

	var totalrewardsClaimed float64

	for i := 1; i <= int(nomination); i++ {
		trc := strconv.Itoa(i)
		trcendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/nominationPools/storage/rewardPools?keys[]=" + trc
		res, err := http.Get(trcendpoint)
		if err != nil {
			fmt.Println("failed to fetch total rewards claimed value", err)
			return 0, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Printf("failed to fetch current total rewards claimed code %d\n", res.StatusCode)
			return 0, err
		}

		var response types.RewardPoolClaimed
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			fmt.Println("Failed to unmarshal total rewards claimed JSON:", err)
			return 0, err
		}
		rewardclaim := response.Value.TotalRewardsClaimed
		rc, err := strconv.ParseFloat(rewardclaim, 64)
		if err != nil {
			fmt.Println("failed to convert rewards claim to int", err)
			return 0, err
		}

		npreward := math.Floor(rc / math.Pow(10, 18))
		totalrewardsClaimed = totalrewardsClaimed + npreward
		return totalrewardsClaimed, nil

	}
	return 0, nil

}
