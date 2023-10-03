package types

type (
	QueryParams map[string]string

	HTTPOptions struct {
		Endpoint    string
		Method      string
		Body        Payload
		QueryParams QueryParams
	}

	Payload struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}

	PingResp struct {
		StatusCode int
		Body       []byte
	}
	Version struct {
		Result struct {
			ClientVersion string `json:"clientVersion"`
		} `json:"result"`
	}
)
