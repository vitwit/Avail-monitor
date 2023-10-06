package monitor

import (
	"fmt"
	"log"
	"os"

	"github.com/vitwit/avail-monitor/config"
)

func GetVersion(cfg *config.Config) {
	version, err := os.ReadFile("config.toml")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(version))
}
