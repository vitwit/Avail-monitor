package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/monitor"
)

type availCollector struct {
	config      *config.Config
	nodeVersion *prometheus.Desc
	chainName   *prometheus.Desc
	// totaltokensIssued *prometheus.Desc
	// nominationPool      *prometheus.Desc
	// bondedToken        *prometheus.Desc
	// bountyProposalCount *prometheus.Desc
	councilMember *prometheus.Desc
	electedMember *prometheus.Desc
}

func NewAvailCollector(cfg *config.Config) *availCollector {
	return &availCollector{
		config: cfg,
		nodeVersion: prometheus.NewDesc(
			"node_version",
			"Node Version Information",
			[]string{"version"}, nil),
		chainName: prometheus.NewDesc(
			"chain",
			"Name of the chain",
			[]string{"chain"}, nil),
		// totaltokensIssued: prometheus.NewDesc(
		// 	"total_tokens_issued",
		// 	"total tokens issued on network",
		// 	[]string{"value"}, nil),
		councilMember: prometheus.NewDesc(
			"council_member_value",
			"council members of the network",
			[]string{"value"}, nil),
		electedMember: prometheus.NewDesc(
			"current_elected_member",
			"elected members of the network",
			[]string{"value"}, nil),
	}
}

func (c *availCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.nodeVersion
	ch <- c.chainName
	// ch <- c.totaltokensIssued
	ch <- c.councilMember
	ch <- c.electedMember

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

	// tokensIssued, err := monitor.FetchTotalTokensIssued(c.config)
	// if err != nil {
	// 	ch <- prometheus.NewInvalidMetric(c.totaltokensIssued, err)
	// } else {
	// 	ch <- prometheus.MustNewConstMetric(c.totaltokensIssued, prometheus.GaugeValue, 1, tokensIssued)
	// }

	councilMem, err := monitor.FetchCouncilMember(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.councilMember, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.councilMember, prometheus.GaugeValue, 1, councilMem)
	}

	electedMem, err := monitor.FetchElectedMember(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.electedMember, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.electedMember, prometheus.GaugeValue, 1, electedMem)
	}

}
