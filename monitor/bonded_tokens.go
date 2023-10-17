package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/vitwit/avail-monitor/config"
// )

// func FetchbondedToken(cfg *config.Config) (string, error) {
// 	btendpoint := cfg.Endpoint.URLEndpoint + "/pallets/staking/storage/erasTotalStake?keys[]="
// 	fmt.Println("bonded token endpoint:", btendpoint)
// 	res, err := http.Get(btendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch bonded token value", err)
// 		return "", err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current bonded token code %d\n", res.StatusCode)
// 		return "", err
// 	}

// 	var result types.bondedTokens
// 	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
// 		fmt.Println("Failed to unmarshal bonded token JSON:", err)
// 		return "", err
// 	}
// 	bonded := result.Value
// 	return bonded, nil

// 	// 	fmt.Printf("bonded: %v\n", bonded)
// 	// 	z, err := strconv.ParseFloat(bonded, 64)
// 	// 	if err != nil {
// 	// 		fmt.Printf("err: %v\n", err)
// 	// 	}
// 	// 	fmt.Printf("z: %v\n", z)
// 	// 	m := math.Floor(z / math.Pow(10, 18))

// 	// bondedToken.WithLabelValues(fmt.Sprintf("%.11e", m)).Set(1.0)
// }
