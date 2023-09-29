package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// prometheus.MustRegister() //need to pass argument
// http.Handle("/metrics", promhttp.Handler())

// Response represents the desired JSON format
type Response struct {
	Version string `json:"version"`
}

func main() {
	// cfg, err := config.ReadFromFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collector := exporter.NewAvailMetric(cfg)

	// prometheus.MustRegister(collector)
	// http.Handle("/metrics", promhttp.Handler()) // exported metrics can be seen in /metrics
	// err = http.ListenAndServe(fmt.Sprintf("%s", cfg.Prometheus.ListenAddress), nil)
	// if err != nil {
	// 	log.Printf("Error while listening on server : %v", err)
	// }

	//WebSocket URL
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
