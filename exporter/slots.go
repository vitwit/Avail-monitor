package exporter

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/monitor"
)

const slotTimeOut = 25 * time.Second

// // scrape time = 25 secs
// // two different scrape times if yes {different set of metrics}
// // all the queries in single ws conn.
var (
	latestBestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "latest_best_block",
		Help: "best block of the network",
	})

	latestFinalizedBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "latest_finalized_block",
		Help: "finalized block of the network",
	})

	timestampOfLatestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "timestamp_of_latest_block",
		Help: "timestamp of the best block in unix milliseconds",
	})

	currentSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_slot",
		Help: "current slot number being queried",
	})

	currentEpoch = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_epoch",
		Help: "current epoch being queried",
	})

	currentEpochStartTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_epoch_start_time",
		Help: "Block height on which the current epoch started",
	})

	currentEpochEndTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_epoch_end_time",
		Help: "Block height on which the current epoch ends",
	})

	totalTokensIssued = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_tokens_issued",
		Help: "total tokens issued on the network",
	})

	// 	totalBondedTokens = prometheus.NewGauge(prometheus.GaugeOpts{
	// 		Name: "total_bonded_tokens",
	// 		Help: "total number of units issued on the network",
	// 	})

	currentEra = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_era",
		Help: "current era",
	})

	bountyProposals = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bounty_proposals",
		Help: "numbet of bounty proposals made on the network",
	})

	nominationPool = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "nomination_pool",
		Help: "number of nomination pools",
	})

	// councilMembers = prometheus.NewGauge(prometheus.GaugeOpts{
	// 	Name: "council_members",
	// 	Help: "council members",
	// })

	totalCouncilProposals = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_council_proposals",
		Help: "number of total council proposals on the network",
	})

	totalPublicProposals = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_public_proposals",
		Help: "number of total public proposals on the network",
	})

	totalPublicReferendums = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_public_referendums",
		Help: "number of total public referendums on the network",
	})

// 	currentElectedMembers = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_elected_members",
// 		Help: "current elected members",
// 	})

//	currentValidators = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "current_validators",
//		Help: "list of current validators",
//	})
)

func init() {

	prometheus.MustRegister(latestBestBlock)
	prometheus.MustRegister(latestFinalizedBlock)
	prometheus.MustRegister(timestampOfLatestBlock)
	prometheus.MustRegister(currentSlot)
	prometheus.MustRegister(currentEpoch)
	prometheus.MustRegister(currentEpochStartTime)
	prometheus.MustRegister(currentEpochEndTime)
	prometheus.MustRegister(totalTokensIssued)
	// 	prometheus.MustRegister(totalBondedTokens)
	prometheus.MustRegister(currentEra)
	prometheus.MustRegister(bountyProposals)
	prometheus.MustRegister(nominationPool)
	// prometheus.MustRegister(councilMembers)
	prometheus.MustRegister(totalCouncilProposals)
	prometheus.MustRegister(totalPublicProposals)
	prometheus.MustRegister(totalPublicReferendums)
	// 	prometheus.MustRegister(currentElectedMembers)
	// 	prometheus.MustRegister(currentValidators)

}

func (c *availCollector) WatchSlots(cfg *config.Config) {
	ticker := time.NewTicker(slotTimeOut)

	for {
		<-ticker.C

		timestamp, err := monitor.FetchTimeStamp(c.config)
		if err != nil {
			log.Printf("Error while getting timestamp: %v", err)
		}
		ts, err := strconv.ParseFloat(timestamp, 64)
		if err != nil {
			log.Printf("Error while converting to float: %v", err)
		}
		timestampOfLatestBlock.Set(ts)

		bestBlock, err := monitor.FetchBestBlock(c.config)
		if err != nil {
			log.Printf("Error while fetching best block: %v", err)
		}
		bb, err := strconv.ParseFloat(bestBlock, 64)
		if err != nil {
			log.Printf("Error while converting best block: %v", err)
		}
		latestBestBlock.Set(bb)

		finalblock, err := monitor.FetchFinalizedBlock(c.config)
		if err != nil {
			log.Printf("Error while fetching finalized block: %v", err)
		}
		fb, err := strconv.ParseFloat(finalblock, 64)
		if err != nil {
			log.Printf("Error while converting finalized block %v", err)
		}
		latestFinalizedBlock.Set(fb)

		slot, err := monitor.FetchCurrentSlot(c.config)
		if err != nil {
			log.Printf("Error while fetching current slot %v", err)
		}
		cs, err := strconv.ParseFloat(slot, 64)
		if err != nil {
			log.Printf("Error while converting current slot: %v", err)
		}
		currentSlot.Set(cs)

		bountyP, err := monitor.FetchBountyProposalCount(c.config)
		if err != nil {
			log.Printf("Error while fetching bounty proposal count: %v", err)
		}
		bp, err := strconv.ParseFloat(bountyP, 64)
		if err != nil {
			log.Printf("Error while converting bounty proposal: %v", err)
		}
		bountyProposals.Set(bp)

		endT, err := monitor.FetchEpochEndTime(c.config)
		if err != nil {
			log.Printf("Error while fetching epoch end time: %v", err)
		}
		eT, err := strconv.ParseFloat(endT, 64)
		if err != nil {
			log.Printf("Error while converting epoch end time: %v", err)
		}
		currentEpochEndTime.Set(eT)

		startT, err := monitor.FetchEpochStartTime(c.config)
		if err != nil {
			log.Printf("Error while fetching epoch start time: %v", err)
		}
		sT, err := strconv.ParseFloat(startT, 64)
		if err != nil {
			log.Printf("Error while converting epoch end time: %v", err)
		}
		currentEpochStartTime.Set(sT)

		cEpoch, err := monitor.FetchEpochIndex(c.config)
		if err != nil {
			log.Printf("Error while fetching current epoch: %v", err)
		}
		cE, err := strconv.ParseFloat(cEpoch, 64)
		if err != nil {
			log.Printf("Error while converting current epoch: %v", err)
		}
		currentEpoch.Set(cE)

		cEra, err := monitor.FetchCurrentEra(c.config)
		if err != nil {
			log.Printf("Error while fetching current era: %v", err)
		}
		ce, err := strconv.ParseFloat(cEra, 64)
		if err != nil {
			log.Printf("Error while converting current era: %v", err)
		}
		currentEra.Set(ce)

		publicProposal, err := monitor.FetchPublicProposalCount(c.config)
		if err != nil {
			log.Printf("Error while fetching public proposal count: %v", err)
		}
		pp, err := strconv.ParseFloat(publicProposal, 64)
		if err != nil {
			log.Printf("Error while converting public proposal count: %v", err)
		}
		totalPublicProposals.Set(pp)

		referendumC, err := monitor.FetchReferendumCount(c.config)
		if err != nil {
			log.Printf("Error while fetching referendum count: %v", err)
		}
		rc, _ := strconv.ParseFloat(referendumC, 64)
		totalPublicReferendums.Set(rc)

		tokensIssued, err := monitor.FetchTotalTokensIssued(c.config)
		if err != nil {
			log.Printf("Error while fetching total tokens issued: %v", err)
		}
		tt, err := strconv.ParseFloat(tokensIssued, 64)
		if err != nil {
			log.Printf("Error while converting total tokens: %v", err)
		}
		abcd := math.Floor(tt / math.Pow(10, 18))
		totalTokensIssued.Set(abcd)

		nPool, err := monitor.FetchNominationPool(c.config)
		if err != nil {
			log.Printf("Error while fetching nomination pool value: %v", err)
		}
		np, err := strconv.ParseFloat(nPool, 64)
		if err != nil {
			log.Printf("Error while converting nomination pool: %v", err)
		}
		nominationPool.Set(np)

		councilproposal, err := monitor.FetchCouncilProposalCount(c.config)
		if err != nil {
			log.Printf("Error while fetching council proposal count: %v", err)
		}
		cpc, err := strconv.ParseFloat(councilproposal, 64)
		if err != nil {
			log.Printf("Error while converting council proposal count: %v", err)
		}
		totalCouncilProposals.Set(cpc)

	}
}
