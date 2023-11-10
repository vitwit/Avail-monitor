package exporter

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/monitor"
)

type availCollector struct {
	config           *config.Config
	nodeVersion      *prometheus.Desc
	chainName        *prometheus.Desc
	councilMember    *prometheus.Desc
	electedMember    *prometheus.Desc
	currentValidator *prometheus.Desc
}

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

		currentValidator: prometheus.NewDesc(
			"avail_monitor_current_validator",
			"current validators of the network",
			[]string{"value"}, nil),
	}
}

func (c *availCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.nodeVersion
	ch <- c.chainName
	ch <- c.councilMember
	ch <- c.electedMember
	ch <- c.currentValidator
}

func (c *availCollector) Collect(ch chan<- prometheus.Metric) {
	version, err := monitor.FetchVersion(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.nodeVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.nodeVersion, prometheus.GaugeValue, 1, version)
	}

	chain, err := monitor.FetchChainID(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.nodeVersion, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.chainName, prometheus.GaugeValue, 1, chain)
	}

	currentVal, err := monitor.FetchCurrentValidators(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.currentValidator, err)
	} else {
		currentvalidator := strings.Join(currentVal, ", ")
		ch <- prometheus.MustNewConstMetric(c.currentValidator, prometheus.GaugeValue, 1, currentvalidator)
	}
}
