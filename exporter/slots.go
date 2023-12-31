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

// scrape time = 25 secs
const slotTimeOut = 25 * time.Second

var (
	latestBestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_chain_best_block",
		Help: "best block of the network",
	})

	timestampOfLatestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_chain_block_timestamp",
		Help: "timestamp of the best block in unix milliseconds",
	})

	currentSlot = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_babe_current_slot",
		Help: "current slot number being queried",
	})

	currentEpoch = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_babe_current_epoch",
		Help: "current epoch being queried",
	})

	currentEpochStartTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_babe_current_epoch_start",
		Help: "Block height on which the current epoch started",
	})

	currentEpochEndTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_babe_current_epoch_end",
		Help: "Block height on which the current epoch ends",
	})

	totalTokensIssued = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_balances_total_tokens_issued",
		Help: "total tokens issued on the network",
	})

	totalBondedTokens = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_staking_total_bonded_tokens",
		Help: "total number of units issued on the network",
	})

	currentEra = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_staking_current_era",
		Help: "current era",
	})

	nominationPool = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_nomination_pools",
		Help: "number of nomination pools",
	})

	totalRewardsDistributed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_total_rewards_distributed",
		Help: "total rewards distributed to validator in era",
	})

	currentStakingRatio = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_current_staking_ratio",
		Help: "current staking ratio",
	})

	totalRewardsClaimed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_total_rewards_claimed",
		Help: "total rewards claimed",
	})

	recordedRewardCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_last_recorded_reward",
		Help: "last recorded reward counter",
	})

	currentValidator = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "avail_monitor_current_validator",
		Help: "total count of current validators",
	})
)

func init() {

	prometheus.MustRegister(latestBestBlock)
	prometheus.MustRegister(timestampOfLatestBlock)
	prometheus.MustRegister(currentSlot)
	prometheus.MustRegister(currentEpoch)
	prometheus.MustRegister(currentEpochStartTime)
	prometheus.MustRegister(currentEpochEndTime)
	prometheus.MustRegister(totalTokensIssued)
	prometheus.MustRegister(totalBondedTokens)
	prometheus.MustRegister(currentEra)
	prometheus.MustRegister(nominationPool)
	prometheus.MustRegister(totalRewardsDistributed)
	prometheus.MustRegister(currentStakingRatio)
	prometheus.MustRegister(totalRewardsClaimed)
	prometheus.MustRegister(recordedRewardCounter)
	prometheus.MustRegister(currentValidator)
}

func (c *availCollector) WatchSlots(cfg *config.Config) {
	ticker := time.NewTicker(slotTimeOut)

	for {
		<-ticker.C

		//get timestamp of latest block of avail network
		timestamp, err := monitor.FetchTimeStamp(c.config)
		if err != nil {
			log.Printf("Error while getting timestamp: %v", err)
		}
		ts, err := strconv.ParseFloat(timestamp, 64)
		if err != nil {
			log.Printf("Error while converting to float: %v", err)
		}
		timestampOfLatestBlock.Set(ts)

		//get best block of the avail network
		bestBlock, err := monitor.FetchBestBlock(c.config)
		if err != nil {
			log.Printf("Error while fetching best block: %v", err)
		}
		bb, err := strconv.ParseFloat(bestBlock, 64)
		if err != nil {
			log.Printf("Error while converting best block: %v", err)
		}
		latestBestBlock.Set(bb)

		//get current slot of the network
		slot, err := monitor.FetchCurrentSlot(c.config)
		if err != nil {
			log.Printf("Error while fetching current slot %v", err)
		}
		cs, err := strconv.ParseFloat(slot, 64)
		if err != nil {
			log.Printf("Error while converting current slot: %v", err)
		}
		currentSlot.Set(cs)

		//get current epoch end time of avail network
		endT, err := monitor.FetchEpochEndTime(c.config)
		if err != nil {
			log.Printf("Error while fetching epoch end time: %v", err)
		}
		eT, err := strconv.ParseFloat(endT, 64)
		if err != nil {
			log.Printf("Error while converting epoch end time: %v", err)
		}
		currentEpochEndTime.Set(eT)

		//get current epoch start time of avail network
		startT, err := monitor.FetchEpochStartTime(c.config)
		if err != nil {
			log.Printf("Error while fetching epoch start time: %v", err)
		}
		sT, err := strconv.ParseFloat(startT, 64)
		if err != nil {
			log.Printf("Error while converting epoch end time: %v", err)
		}
		currentEpochStartTime.Set(sT)

		//get current epoch of avail network
		cEpoch, err := monitor.FetchEpochIndex(c.config)
		if err != nil {
			log.Printf("Error while fetching current epoch: %v", err)
		}
		cE, err := strconv.ParseFloat(cEpoch, 64)
		if err != nil {
			log.Printf("Error while converting current epoch: %v", err)
		}
		currentEpoch.Set(cE)

		//get current era of avail network
		cEra, err := monitor.FetchCurrentEra(c.config)
		if err != nil {
			log.Printf("Error while fetching current era: %v", err)
		}
		ce, err := strconv.ParseFloat(cEra, 64)
		if err != nil {
			log.Printf("Error while converting current era: %v", err)
		}
		currentEra.Set(ce)

		//get total rewards distributed on avail network
		totalrewardsdist, err := monitor.FetchTotalRewardsDistributed(c.config)
		if err != nil {
			log.Printf("Error while fetching total rewards distributed to a val in era: %v", err)
		}
		trd, err := strconv.ParseFloat(totalrewardsdist, 64)
		if err != nil {
			log.Printf("Error while converting total rewards distributed to a val in era: %v", err)
		}
		totalRewardsDistributed.Set(trd)

		//get total rewards claimed on avail network
		totalrewardsclaimed, err := monitor.FetchTotalRewardsClaimed(c.config)
		if err != nil {
			log.Printf("Error while fetching total rewards claimed: %v", err)
		}
		totalRewardsClaimed.Set(totalrewardsclaimed)

		//get nominator pool reward on avail network
		poolReward, err := monitor.FetchNominatorPoolRewards(c.config)
		if err != nil {
			log.Printf("Error while fetching nominator pool rewards: %v", err)
		}
		recordedRewardCounter.Set(poolReward)

		//get total tokens isssued on the avail network
		tokensIssued, err := monitor.FetchTotalTokensIssued(c.config)
		if err != nil {
			log.Printf("Error while fetching total tokens issued: %v", err)
		}
		tt, err := strconv.ParseFloat(tokensIssued, 64)
		if err != nil {
			log.Printf("Error while converting total tokens: %v", err)
		}
		tti := math.Floor(tt / math.Pow(10, 18))
		totalTokensIssued.Set(tti)

		//get total bonded tokens on the network
		bondedtokens, err := monitor.FetchBondedToken(c.config)
		if err != nil {
			log.Printf("Error while fetching bonded tokens: %v", err)
		}
		bt, err := strconv.ParseFloat(bondedtokens, 64)
		if err != nil {
			log.Printf("Error while converting bonded tokens: %v", err)
		}
		bondedtokenstotal := math.Floor(bt / math.Pow(10, 18))
		totalBondedTokens.Set(bondedtokenstotal)

		//get nomination pool on avail network
		nPool, err := monitor.FetchNominationPool(c.config)
		if err != nil {
			log.Printf("Error while fetching nomination pool value: %v", err)
		}
		np, err := strconv.ParseFloat(nPool, 64)
		if err != nil {
			log.Printf("Error while converting nomination pool: %v", err)
		}
		nominationPool.Set(np)

		//get current staking ratio on avail network
		currentSR, err := monitor.FetchCurrentStakingRatio(c.config)
		if err != nil {
			log.Printf("Error while fetching current staking ratio value: %v", err)
		}
		currentStakingRatio.Set(currentSR)

		//get no. of current validator on avail network
		currentVal, err := monitor.FetchCurrentValidators(c.config)
		if err != nil {
			log.Printf("Error while getting count of current validators: %v", err)
		}
		currentValidator.Set(float64(currentVal))
	}
}
