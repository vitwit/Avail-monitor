package exporter

// import (
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// )

// const slotSchedule = 25 * time.Second

// // scrape time = 25 secs
// // two different scrape times if yes {different set of metrics}
// // all the queries in single ws conn.
// var (
// 	clientVersion = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "avail_node_version",
// 		Help: "node version of avail",
// 	})

// 	chainID = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "chain_id",
// 		Help: "chain id of the network",
// 	})

// 	latestBestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "latest_best_block",
// 		Help: "best block of the network",
// 	})

// 	latestFinalizedBlock = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "latest_finalized_block",
// 		Help: "finalized block of the network",
// 	})

// 	timestampOfLatestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "timestamp_of_latest_block",
// 		Help: "timestamp of the best block in unix milliseconds",
// 	})

// 	currentSlot = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_slot",
// 		Help: "current slot number being queried",
// 	})

// 	currentEpoch = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_epoch",
// 		Help: "current epoch being queried",
// 	})

// 	currentEpochStartTime = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_epoch_start_time",
// 		Help: "Block height on which the current epoch started",
// 	})

// 	currentEpochEndTime = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_epoch_end_time",
// 		Help: "Block height on which the current epoch ends",
// 	})

// 	totalTokensIssued = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "total_tokens_issued",
// 		Help: "total tokens issued on the network",
// 	})

// 	totalBondedTokens = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "total_bonded_tokens",
// 		Help: "total number of units issued on the network",
// 	})

// 	currentEra = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_era",
// 		Help: "current era",
// 	})

// 	bountyProposals = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "bounty_proposals",
// 		Help: "numbet of bounty proposals made on the network",
// 	})

// 	councilMembers = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "council_members",
// 		Help: "council members",
// 	})

// 	totalCouncilProposals = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "total_council_proposals",
// 		Help: "number of total council proposals on the network",
// 	})

// 	totalPublicProposals = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "total_public_proposals",
// 		Help: "number of total public proposals on the network",
// 	})

// 	totalPublicReferendums = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "total_public_referendums",
// 		Help: "number of total public referendums on the network",
// 	})

// 	currentElectedMembers = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_elected_members",
// 		Help: "current elected members",
// 	})

// 	currentValidators = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_validators",
// 		Help: "list of current validators",
// 	})
// )

// func init() {
// 	//prometheus.MustRegister(clientVersion)
// 	prometheus.MustRegister(chainID)
// 	prometheus.MustRegister(latestBestBlock)
// 	prometheus.MustRegister(latestFinalizedBlock)
// 	prometheus.MustRegister(timestampOfLatestBlock)
// 	prometheus.MustRegister(currentSlot)
// 	prometheus.MustRegister(currentEpoch)
// 	prometheus.MustRegister(currentEpochStartTime)
// 	prometheus.MustRegister(currentEpochEndTime)
// 	prometheus.MustRegister(totalTokensIssued)
// 	prometheus.MustRegister(totalBondedTokens)
// 	prometheus.MustRegister(currentEra)
// 	prometheus.MustRegister(bountyProposals)
// 	prometheus.MustRegister(councilMembers)
// 	prometheus.MustRegister(totalCouncilProposals)
// 	prometheus.MustRegister(totalPublicProposals)
// 	prometheus.MustRegister(totalPublicReferendums)
// 	prometheus.MustRegister(currentElectedMembers)
// 	prometheus.MustRegister(currentValidators)

// }
