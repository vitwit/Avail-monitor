package monitor

import (
	"fmt"

	"github.com/vitwit/avail-monitor/types"

	"github.com/vitwit/avail-monitor/config"
)

func fetchCurrentSlot(cfg *config.Config) (types.CurrentSlot, error) {
	var result types.CurrentSlot
	fmt.Println(result)
	return result, nil
}
