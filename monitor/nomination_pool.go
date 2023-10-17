package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchNominationPool(cfg *config.Config) (string, error) {
	poolendpoint := cfg.Endpoint.URLEndpoint + "/pallets/nominationPools/storage/counterForBondedPools"
	fmt.Printf("epochindex enddpoint: %v\n", poolendpoint)
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
		fmt.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	nominationpool := response.Value
	return nominationpool, nil

	// np, _ := strconv.ParseFloat(nominationpool, 64)
	// fmt.Printf("np................ %v\n", np)
	// nominationPool.Set(np)

}
