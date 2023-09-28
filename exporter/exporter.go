package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/monitor"
)

type availMetric struct {
	config       *config.Config
	availVersion *prometheus.Desc
}

func NewAvailMetric(cfg *config.Config) *availMetric {
	return &availMetric{
		config: cfg,
		availVersion: prometheus.NewDesc(
			"avail_node_version",
			"node version of avail",
			[]string{"version"}, nil),
	}
}

func (c *availMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.availVersion
}

func (c *availMetric) Collect(ch chan<- prometheus.Metric) {
	ver, err := monitor.GetVersion(c.config)
	if ver.Result.version != "" {
		ch <- prometheus.MustNewConstMetric(c.availVersion, prometheus.GaugeValue, 1, ver.Result.version)
	}
}
