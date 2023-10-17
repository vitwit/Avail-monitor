package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitwit/avail-monitor/config"
	"github.com/vitwit/avail-monitor/exporter"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	collector := exporter.NewAvailCollector(cfg)

	go collector.WatchSlots(cfg)

	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler()) // exported metrics can be seen in /metrics
	err = http.ListenAndServe(fmt.Sprintf("%s", cfg.Prometheus.ListenAddress), nil)
	if err != nil {
		log.Printf("Error while listening on server : %v", err)
	}

}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"math"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"github.com/vitwit/avail-monitor/config"
// 	"github.com/vitwit/avail-monitor/monitor"
// 	"github.com/vitwit/avail-monitor/types"
// )

// const (
// 	apiEndpoint  = "http://64.227.177.52:8080"
// 	metricsPort  = 9000
// 	timeInterval = 25 * time.Second
// )

// var (
// 	// nodeVersion = prometheus.NewGaugeVec(
// 	// 	prometheus.GaugeOpts{
// 	// 		Name: "node_version",
// 	// 		Help: "Node Version Information",
// 	// 	},
// 	// 	[]string{"version"},
// 	// )
// 	chainName = prometheus.NewGaugeVec(
// 		prometheus.GaugeOpts{
// 			Name: "chain",
// 			Help: "Name of the chain",
// 		},
// 		[]string{"chain"},
// 	)
// 	currentSlot = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_slot_value",
// 		Help: "current slot value of avail",
// 	},
// 	)
// 	epochIndex = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_epoch_number",
// 		Help: "current epoch number of avail",
// 	},
// 	)
// 	timeStamp = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "timestamp_of_latest_block",
// 		Help: "timestamp of latest block",
// 	},
// 	// []string{"now"},
// 	)
// 	bestBlock = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "latest_best_block",
// 		Help: "latest best block",
// 	},
// 	)
// 	finalizedBlock = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "finalized_block",
// 		Help: "finalized block of the network",
// 	},
// 	)
// 	epochstartTime = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "epoch_start_time",
// 		Help: "epoch start time of network",
// 	},
// 	)
// 	epochendTime = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "epoch_end_time",
// 		Help: "epoch end time of network",
// 	},
// 	)
// 	totaltokensIssued = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "total_tokens_issued",
// 		Help: "total tokens issued on network",
// 	}, []string{"value"})
// 	nominationPool = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "nomination_pool",
// 		Help: "number of nomination pools",
// 	})
// 	currentEra = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "current_era_value",
// 		Help: "current era",
// 	})
// 	bondedToken = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "total_bonded_token_value",
// 		Help: "bonded token value",
// 	}, []string{"value"})
// 	proposalCount = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "proposal_count_value",
// 		Help: "total proposal count value",
// 	})
// 	referendumCount = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "referendum_count_value",
// 		Help: "total referendum count value",
// 	})
// 	publicProposalCount = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "public_proposal_count_value",
// 		Help: "total public proposal count value",
// 	})
// 	bountyProposalCount = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "bounty_proposal_count_value",
// 		Help: "total bounty proposal count value",
// 	})
// 	councilMember = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "council_member_value",
// 		Help: "total council member value",
// 	}, []string{"value"})
// 	electedMember = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "current_elected_member",
// 		Help: "current elected member value",
// 	}, []string{"value"})
// )

// func fetchDataAndSetMetric() {
// 	endpoint := apiEndpoint + "/node/version"
// 	fmt.Printf("apiEndpoint: %v\n", endpoint)
// 	resp, err := http.Get(endpoint)
// 	if err != nil {
// 		fmt.Println("Failed to fetch data:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Failed to fetch data. Status code: %d\n", resp.StatusCode)
// 		return
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Failed to read response body:", err)
// 		return
// 	}

// 	var data map[string]string
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	// version, found := data["clientVersion"]
// 	// if !found {
// 	// 	fmt.Println("Version not found in response")
// 	// 	return
// 	// }

// 	chain, found := data["chain"]
// 	if !found {
// 		fmt.Println("Chain not found in response")
// 		return
// 	}

// 	// nodeVersion.WithLabelValues(version).Set(1)
// 	chainName.WithLabelValues(chain).Set(1)
// 	fmt.Printf("chain: %s\n", chain)
// 	// fmt.Printf("Node Version: %s\n", version)
// }

// func fetchCurrentSlot() {
// 	finalendpoint := apiEndpoint + "/pallets/babe/storage/currentSlot"
// 	fmt.Printf("currentSlot: %v\n", finalendpoint)
// 	resp, err := http.Get(finalendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch finalied block", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch finalzed code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CurrentSlot
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	value := response.Value

// 	v, _ := strconv.ParseFloat(value, 64)
// 	fmt.Println("value here....", v)
// 	currentSlot.Set(v)
// 	fmt.Printf("current slot Value: %s\n", value)

// }

// func fetchEpochIndex() {
// 	epochendpoint := apiEndpoint + "/pallets/babe/storage/epochIndex"
// 	fmt.Printf("epochindex enddpoint: %v\n", epochendpoint)
// 	resp, err := http.Get(epochendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochIndex
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	value := response.Value
// 	e, _ := strconv.ParseFloat(value, 64)
// 	epochIndex.Set(e)
// 	fmt.Printf("epoch index value: %s\n", value)

// }

// func fetchTimeStamp() {
// 	tsendpoint := apiEndpoint + "/blocks/head"
// 	resp, err := http.Get(tsendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch timestamp", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.TimeStamp
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}
// 	epochtime := response.Extrinsics[0].Args.Now
// 	fmt.Printf("epochtime: %v\n", epochtime)
// 	ts, _ := strconv.ParseFloat(epochtime, 64)
// 	timeStamp.Set(ts) // Export as seconds

// 	fmt.Printf("Fetched timestamp: %s\n", epochtime)
// }

// func fetchBestBlock() {
// 	blockendpoint := apiEndpoint + "/blocks/head"
// 	fmt.Printf("epochindex enddpoint: %v\n", blockendpoint)
// 	resp, err := http.Get(blockendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.BestBlock
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	block := response.Number
// 	b, _ := strconv.ParseFloat(block, 64)
// 	bestBlock.Set(b)
// }

// func fetchFinalizedBlock() {
// 	finalizedendpoint := apiEndpoint + "/blocks/head"
// 	resp, err := http.Get(finalizedendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch index", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch index code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.FinalizedBlock
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	finalizedblock := response.Hash
// 	h, _ := strconv.ParseFloat(finalizedblock, 64)
// 	fmt.Printf("finalized block ***********%v\n", h)
// 	finalizedBlock.Set(h)

// }

// func fetchEpochStartTime() {
// 	startendpoint := apiEndpoint + "/pallets/babe/storage/epochStart"
// 	fmt.Printf("epoch start time: %v\n", startendpoint)
// 	resp, err := http.Get(startendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch start time", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch start time code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochStartTime
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	startTime := response.Value[0]
// 	fmt.Println(startTime)
// 	st, _ := strconv.ParseFloat(startTime, 64)
// 	epochstartTime.Set(st)

// }

// func fetchEpochEndTime() {
// 	epochendpoint := apiEndpoint + "/pallets/babe/storage/epochStart"
// 	fmt.Printf("epoch end time: %v\n", epochendpoint)
// 	resp, err := http.Get(epochendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch epoch end time", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch end time code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.EpochEndTime
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	endTime := response.Value[1]
// 	fmt.Println(endTime)
// 	et, _ := strconv.ParseFloat(endTime, 64)
// 	epochendTime.Set(et)

// }

// func fetchTotalTokensIssued() {
// 	tokenendpoint := apiEndpoint + "/pallets/balances/storage/totalIssuance"
// 	resp, err := http.Get(tokenendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch total token issuance", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch total token issuance code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.TotalTokensIssued
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	totalTokens := response.Value
// 	tt, _ := strconv.ParseFloat(totalTokens, 64)
// 	abcd := math.Floor(tt / math.Pow(10, 18))
// 	fmt.Printf("abcd: %v\n", abcd)

// 	// ttI := tt / 1e18 //wrong conversion.. consider later..
// 	//totaltokensIssued.Set(abcd)
// 	totaltokensIssued.WithLabelValues(fmt.Sprintf("%.11e", abcd)).Set(1)

// }

// func fetchNominationPool() {
// 	poolendpoint := apiEndpoint + "/pallets/nominationPools/storage/counterForBondedPools"
// 	fmt.Printf("epochindex enddpoint: %v\n", poolendpoint)
// 	resp, err := http.Get(poolendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch nomination pools", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch number of nomination pools %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.NominationPool
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	nominationpool := response.Value
// 	np, _ := strconv.ParseFloat(nominationpool, 64)
// 	fmt.Printf("np................ %v\n", np)
// 	nominationPool.Set(np)

// }

// func fetchCurrentEra() {
// 	eraendpoint := apiEndpoint + "/pallets/staking/storage/currentEra"
// 	fmt.Printf("currentSlot: %v\n", eraendpoint)
// 	resp, err := http.Get(eraendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch current era value", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current era value %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CurrentEra
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}
// 	value := response.Value
// 	e, _ := strconv.ParseFloat(value, 64)
// 	currentEra.Set(e)

// 	//***** bonded-tokens
// 	btendpoint := apiEndpoint + "/pallets/staking/storage/erasTotalStake?keys[]=" + response.Value
// 	fmt.Println("bonded token endpoint:", btendpoint)
// 	res, err := http.Get(btendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch bonded token value", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current bonded token code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var result types.bondedTokens
// 	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
// 		fmt.Println("Failed to unmarshal bonded token JSON:", err)
// 		return
// 	}
// 	bonded := result.Value
// 	fmt.Printf("bonded: %v\n", bonded)
// 	z, err := strconv.ParseFloat(bonded, 64)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("z: %v\n", z)
// 	m := math.Floor(z / math.Pow(10, 18))

// 	bondedToken.WithLabelValues(fmt.Sprintf("%.11e", m)).Set(1.0)

// }

// func fetchProposalCount() {
// 	pcendpoint := apiEndpoint + "/pallets/council/storage/proposalCount"
// 	fmt.Printf("proposalendpoint: %v\n", pcendpoint)
// 	resp, err := http.Get(pcendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch proposal count", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch proposal count %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.ProposalCount
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	value := response.Value
// 	pc, _ := strconv.ParseFloat(value, 64)
// 	proposalCount.Set(pc)

// }

// func fetchReferendumCount() {
// 	rcendpoint := apiEndpoint + "/pallets/democracy/storage/referendumCount"
// 	resp, err := http.Get(rcendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch referendum count", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch referendum count %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.ReferendumCount
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	referendum := response.Value
// 	rc, _ := strconv.ParseFloat(referendum, 64)
// 	referendumCount.Set(rc)
// }

// func fetchPublicProposalCount() {
// 	ppcendpoint := apiEndpoint + "/pallets/democracy/storage/publicPropCount"
// 	resp, err := http.Get(ppcendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch public proposal count", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch public proposal count%d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.PublicProposalCount
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	publicpc := response.Value
// 	ppc, _ := strconv.ParseFloat(publicpc, 64)
// 	publicProposalCount.Set(ppc)
// }

// func fetchBountyProposalCount() {
// 	bpcendpoint := apiEndpoint + "/pallets/bounties/storage/bountyCount"
// 	resp, err := http.Get(bpcendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch bounty proposal count value", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch bounty proposal count value %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.BountyProposalCount
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	bountypc := response.Value
// 	bpc, _ := strconv.ParseFloat(bountypc, 64)
// 	bountyProposalCount.Set(bpc)
// }

// func fetchCouncilMember() {
// 	cmendpoint := apiEndpoint + "/pallets/council/storage/members"
// 	resp, err := http.Get(cmendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch nomination pools", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch epoch number of nomination pools %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.CouncilMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	councilmem := response.Value[0]
// 	// fmt.Println(councilmem)
// 	// cm, _ := strconv.ParseFloat(councilmem, 32)
// 	// fmt.Printf("cm prom metric-------- %v\n", cm)
// 	councilMember.WithLabelValues(councilmem).Set(1)

// }

// func fetchElectedMember() {
// 	cemendpoint := apiEndpoint + "/pallets/elections/storage/members"
// 	resp, err := http.Get(cemendpoint)
// 	if err != nil {
// 		fmt.Println("failed to fetch current elected member", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("failed to fetch current elected member code %d\n", resp.StatusCode)
// 		return
// 	}

// 	var response types.ElectedMembers
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		fmt.Println("Failed to unmarshal JSON:", err)
// 		return
// 	}

// 	elecmem := response.Value[0].Who
// 	fmt.Printf("elecmem: %v\n", elecmem)
// 	// fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------", elecmem)
// 	// cem, err := strconv.ParseFloat(elecmem, 64)
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }
// 	// fmt.Println("current elected member value------------------------------------------->", cem)
// 	electedMember.WithLabelValues(elecmem).Set(1)
// }

// func main() {
// 	monitor.FetchVersion(&config.Config{})
// 	ticker := time.NewTicker(1 * time.Second)

// 	// version, err := os.ReadFile("config.toml")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// prometheus.MustRegister(nodeVersion)
// 	prometheus.MustRegister(chainName)
// 	prometheus.MustRegister(currentSlot)
// 	prometheus.MustRegister(epochIndex)
// 	prometheus.MustRegister(timeStamp)
// 	prometheus.MustRegister(bestBlock)
// 	prometheus.MustRegister(finalizedBlock)
// 	prometheus.MustRegister(epochstartTime)
// 	prometheus.MustRegister(epochendTime)
// 	prometheus.MustRegister(totaltokensIssued)
// 	prometheus.MustRegister(nominationPool)
// 	prometheus.MustRegister(currentEra)
// 	prometheus.MustRegister(proposalCount)
// 	prometheus.MustRegister(referendumCount)
// 	prometheus.MustRegister(publicProposalCount)
// 	prometheus.MustRegister(bountyProposalCount)
// 	prometheus.MustRegister(councilMember)
// 	prometheus.MustRegister(electedMember)
// 	prometheus.MustRegister(bondedToken)

// 	go func() {
// 		http.Handle("/metrics", promhttp.Handler())
// 		fmt.Printf("Starting Prometheus server on :%d\n", metricsPort)
// 		http.ListenAndServe(fmt.Sprintf(":%d", metricsPort), nil)
// 	}()

// 	for {
// 		<-ticker.C
// 		fetchDataAndSetMetric()
// 		fetchCurrentSlot()
// 		fetchEpochIndex()
// 		fetchTimeStamp()
// 		fetchBestBlock()
// 		fetchFinalizedBlock()
// 		fetchEpochStartTime()
// 		fetchEpochEndTime()
// 		fetchTotalTokensIssued()
// 		fetchNominationPool()
// 		fetchCurrentEra()
// 		fetchProposalCount()
// 		fetchReferendumCount()
// 		fetchPublicProposalCount()
// 		fetchBountyProposalCount()
// 		fetchCouncilMember()
// 		fetchElectedMember()
// 		time.Sleep(timeInterval)
// 	}
// }
