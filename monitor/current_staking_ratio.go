package monitor

import (
	"fmt"
	"strconv"

	"github.com/vitwit/avail-monitor/config"
)

// FetchCurrentStakingRatio displays the ratio of
// ratio of total bonded tokens to total tokens issued on the network and
// multiplying the result by 100
func FetchCurrentStakingRatio(cfg *config.Config) (float64, error) {
	bondedTokens, err := FetchBondedToken(cfg)
	if err != nil {
		return 0, err
	}

	totalTokensIssued, err := FetchTotalTokensIssued(cfg)
	if err != nil {
		return 0, err
	}

	bondedTokenscsr, err := strconv.ParseFloat(bondedTokens, 64)
	if err != nil {
		fmt.Println("Failed to convert bonded tokens to float:", err)
		return 0, err
	}

	totalTokensIssuedcsr, err := strconv.ParseFloat(totalTokensIssued, 64)
	if err != nil {
		fmt.Println("Failed to convert total tokens issued to float:", err)
		return 0, err
	}

	if totalTokensIssuedcsr == 0 {
		fmt.Println("Total tokens issued is 0, cannot calculate percentage")
		return 0, err
	}

	curentstakingratio := (bondedTokenscsr / totalTokensIssuedcsr) * 100
	return curentstakingratio, nil
}
