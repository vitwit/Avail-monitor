package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	//"github.com/vitwit/avail-monitor/types"
)

type availMetric struct {
	config        *config.Config
	clientVersion *prometheus.Desc
}

func NewAvailMetric(cfg *config.Config) *availMetric {
	return &availMetric{
		config: cfg,
		clientVersion: prometheus.NewDesc(
			"avail_node_version",
			"node version of avail",
			[]string{"avail_node_version"}, nil),
	}
}

func (c *availMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.clientVersion
}

// func (c *availMetric) Collect(ch chan<- prometheus.Metric) {
// 	// fmt.Printf("\"versio\": %v\n", "version")
// 	version, err := monitor.GetVersion(c.config)
// 	// fmt.Printf("version: %v\n", version)
// 	if version.Result.ClientVersion != "" {
// 		ch <- prometheus.MustNewConstMetric(c.clientVersion, prometheus.GaugeValue, 1, version.Result.ClientVersion)
// 	}
// 	if err != nil {
// 		log.Printf("failed to fetch version %s", err)
// 	}
// }
