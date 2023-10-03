package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/types"
)

func GetVersion(cfg *config.Config) (types.Version, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.URLEndpoint,
		Method:   http.MethodGet,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "clientVersion", ID: 1},
	}

	var result types.Version
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	return result, nil
}
