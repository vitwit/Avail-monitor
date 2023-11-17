package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func FetchCouncilMember(cfg *config.Config) (string, error) {
// 	cmendpoint := cfg.RPC_Endpoint.URLEndpoint + "/pallets/council/storage/members"
// 	resp, err := http.Get(cmendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch council member value", err)
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch status code of council member %d\n", resp.StatusCode)
// 		return "", err
// 	}

// 	var response types.CouncilMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to fetch data for council members from endpoint:", err)
// 		return "", err
// 	}

// 	councilmem := response.Value[0]
// 	return councilmem, nil
// }
