package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func FetchTotalTokensIssued(cfg *config.Config) (string, error) {
	tokenendpoint := cfg.Endpoint.URLEndpoint + "/pallets/balances/storage/totalIssuance"
	resp, err := http.Get(tokenendpoint)
	if err != nil {
		fmt.Println("failed to fetch total token issuance", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("failed to fetch total token issuance code %d\n", resp.StatusCode)
		return "", err
	}

	var response types.TotalTokensIssued
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	totalTokens := response.Value
	return totalTokens, nil

	// tt, _ := strconv.ParseFloat(totalTokens, 64)
	// abcd := math.Floor(tt / math.Pow(10, 18))
	// fmt.Printf("abcd: %v\n", abcd)

	// ttI := tt / 1e18 //wrong conversion.. consider later..
	//totaltokensIssued.Set(abcd)
	// totaltokensIssued.WithLabelValues(fmt.Sprintf("%.11e", abcd)).Set(1)

}
