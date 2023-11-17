package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/monitor"
)

// availCollector represents the set of avail metrics
type availCollector struct {
	config      *config.Config
	nodeVersion *prometheus.Desc
	chainName   *prometheus.Desc
}

// /NewAvailCollector exports availCollector metrics to prometheus
func NewAvailCollector(cfg *config.Config) *availCollector {
	return &availCollector{
		config: cfg,
		nodeVersion: prometheus.NewDesc(
			"avail_monitor_chain_node_version",
			"Node Version Information",
			[]string{"version"}, nil),

		chainName: prometheus.NewDesc(
			"avail_monitor_chain_name",
			"Name of the chain",
			[]string{"chain"}, nil),
	}
}

func (c *availCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.nodeVersion
	ch <- c.chainName
}

func (c *availCollector) Collect(ch chan<- prometheus.Metric) {
	//get version of avail network
	version, err := monitor.FetchVersion(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.nodeVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.nodeVersion, prometheus.GaugeValue, 1, version)
	}

	//get chain ID of avail network
	chain, err := monitor.FetchChainID(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.nodeVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.chainName, prometheus.GaugeValue, 1, chain)
	}
}
