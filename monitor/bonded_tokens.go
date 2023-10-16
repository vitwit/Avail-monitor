package monitor

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"net/http"
// 	"strconv"

// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/types"
// )

// func fetchBoundedToken(cfg *config.Config) {
// 	btendpoint := cfg.URLEndpoint + "/pallets/staking/storage/erasTotalStake?keys[]=" + response.Value
// 	fmt.Println("bounded token endpoint:", btendpoint)
// 	res, err := http.Get(btendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch bounded token value", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current bounded token code %d\n", res.StatusCode)
// 		return
// 	}

// 	var result types.BoundedTokens
// 	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
// 		fmt.Println("Failed to unmarshal bounded token JSON:", err)
// 		return
// 	}
// 	bounded := result.Value
// 	fmt.Printf("bounded: %v\n", bounded)
// 	z, err := strconv.ParseFloat(bounded, 64)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("z: %v\n", z)
// 	m := math.Floor(z / math.Pow(10, 18))

// 	boundedToken.WithLabelValues(fmt.Sprintf("%.11e", m)).Set(1.0)
// }
