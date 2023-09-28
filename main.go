// package main

// import (
// 	"net/http"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// func main() {

// 	prometheus.MustRegister() //need to pass argument
// 	http.Handle("/metrics", promhttp.Handler())
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Response represents the desired JSON format
type Response struct {
	Version string `json:"version"`
}

func main() {
	// WebSocket URL
	wsURL := "wss://kate.avail.tools/ws"

	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer conn.Close()

	// Prepare the JSON-RPC request
	request := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "system_version",
	}

	// Send the request
	err = conn.WriteJSON(request)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}

	// Read the response
	var response map[string]interface{}
	err = conn.ReadJSON(&response)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Extract the "result" field
	result, ok := response["result"].(string)
	if !ok {
		log.Fatal("Response does not contain a valid 'result' field")
	}

	// Create the desired response format
	formattedResponse := Response{
		Version: result,
	}

	// Marshal the formatted response into JSON
	jsonResponse, err := json.Marshal(formattedResponse)
	if err != nil {
		log.Fatal("Error marshaling response:", err)
	}

	fmt.Println(string(jsonResponse))
}
