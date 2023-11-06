package types

type CurrentSlot struct {
	Value string `json:"value"`
}

type EpochIndex struct {
	Value string `json:"value"`
}

type TimeStamp struct {
	Extrinsics []struct {
		Args struct {
			Now string `json:"now"`
		} `json:"args"`
	} `json:"extrinsics"`
}

type BestBlock struct {
	Number string `json:"number"`
}

type FinalizedBlock struct {
	Hash string `json:"hash"`
}

type EpochStartTime struct {
	Value []string `json:"value"`
}

type EpochEndTime struct {
	Value []string `json:"value"`
}

type TotalTokensIssued struct {
	Value string `json:"value"`
}

type NominationPool struct {
	Value string `json:"value"`
}

type CurrentEra struct {
	Value string `json:"value"`
}

type ProposalCount struct {
	Value string `json:"value"`
}

type ReferendumCount struct {
	Value string `json:"value"`
}

type PublicProposalCount struct {
	Value string `json:"value"`
}

type BountyProposalCount struct {
	Value string `json:"value"`
}

type CouncilMembers struct {
	Value []string `json:"value"`
}

type CouncilProposals struct {
	Value string `json:"value"`
}

type ElectedMembers struct {
	Value []struct {
		Who string `json:"who"`
	} `json:"value"`
}

type BondedTokens struct {
	Value string `json:"value"`
}

type CurrentValidators struct {
	Value []string `json:"value"`
}

type TotalRewardsDistributed struct {
	Value struct {
		Total string `json:"total"`
	} `json:"value"`
}
