package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		campaign1      = sample.Campaign(0)
		campaign2      = sample.Campaign(1)
		sharesVesting1 = sample.Shares()
		sharesVesting2 = sample.Shares()
		shares1        = sample.Shares()
		shares2        = sample.Shares()
		shares3        = sample.Shares()
		shares4        = sample.Shares()
	)
	sharesCampaign1 := types.IncreaseShares(shares1, shares2)
	sharesCampaign1 = types.IncreaseShares(sharesCampaign1, sharesVesting1)
	campaign1.AllocatedShares = sharesCampaign1
	campaign1.TotalShares = sharesCampaign1
	campaign1.DynamicShares = true

	sharesCampaign2 := types.IncreaseShares(shares3, shares4)
	sharesCampaign2 = types.IncreaseShares(sharesCampaign2, sharesVesting2)
	campaign2.AllocatedShares = sharesCampaign2
	campaign2.TotalShares = sharesCampaign2
	campaign2.DynamicShares = true

	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: campaign1.Id,
					},
					{
						CampaignID: campaign2.Id,
					},
				},
				CampaignList: []types.Campaign{
					campaign1,
					campaign2,
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: campaign1.Id,
						Address:    sample.Address(),
						Shares:     shares1,
					},
					{
						CampaignID: campaign2.Id,
						Address:    sample.Address(),
						Shares:     shares3,
					},
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID:     campaign1.Id,
						Address:        sample.Address(),
						StartingShares: shares2,
						VestingOptions: *types.NewShareDelayedVesting(sharesVesting1, time.Now().Unix()),
					},
					{
						CampaignID:     campaign2.Id,
						Address:        sample.Address(),
						StartingShares: shares4,
						VestingOptions: *types.NewShareDelayedVesting(sharesVesting2, time.Now().Unix()),
					},
				},
			},
			valid: true,
		},
		{
			desc: "non existing campaign for mainnet vesting account",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 1,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(1),
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(0, sample.Address()),
					sample.MainnetAccount(1, sample.Address()),
				},
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: 33333,
						Address:    "33333",
					},
					{
						CampaignID: 9999,
						Address:    "9999",
					},
				},
			},
			valid: false,
		},
		{
			desc: "non existing campaign for mainnet account",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 1,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(1),
				},
				CampaignCounter: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(330, sample.Address()),
					sample.MainnetAccount(434, sample.Address()),
				},
			},
			valid: false,
		},
		{
			desc: "non existing campaign for chains",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 2,
					},
					{
						CampaignID: 4,
					},
				},
				CampaignList: []types.Campaign{
					sample.Campaign(99),
					sample.Campaign(88),
				},
				CampaignCounter: 100,
			},
			valid: false,
		},
		{
			desc: "duplicated campaignChains",
			genState: &types.GenesisState{
				CampaignChainsList: []types.CampaignChains{
					{
						CampaignID: 0,
					},
					{
						CampaignID: 0,
					},
				},
			},
		},
		{
			desc: "duplicated mainnetVestingAccount",
			genState: &types.GenesisState{
				MainnetVestingAccountList: []types.MainnetVestingAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
					{
						CampaignID: 0,
						Address:    "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(0),
					sample.Campaign(0),
				},
				CampaignCounter: 2,
			},
			valid: false,
		},
		{
			desc: "invalid campaign count",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					sample.Campaign(1),
				},
				CampaignCounter: 0,
			},
			valid: false,
		},
		{
			desc: "invalid campaign",
			genState: &types.GenesisState{
				CampaignList: []types.Campaign{
					types.NewCampaign(0, invalidCampaignName, sample.Uint64(), sample.Coins(), false),
				},
				CampaignCounter: 1,
			},
			valid: false,
		},
		{
			desc: "duplicated mainnetAccount",
			genState: &types.GenesisState{
				MainnetAccountList: []types.MainnetAccount{
					{
						CampaignID: 0,
						Address:    "0",
					},
					{
						CampaignID: 0,
						Address:    "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			campaignIDMap := make(map[uint64]types.Shares)
			for _, elem := range tc.genState.CampaignList {
				campaignIDMap[elem.Id] = elem.AllocatedShares
			}
			shares := make(map[uint64]types.Shares)

			for _, acc := range tc.genState.MainnetAccountList {
				// check if the campaign exists for mainnet accounts
				_, ok := campaignIDMap[acc.CampaignID]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.CampaignID]; !ok {
					shares[acc.CampaignID] = types.EmptyShares()
				}
				shares[acc.CampaignID] = types.IncreaseShares(
					shares[acc.CampaignID],
					acc.Shares,
				)
			}

			for _, acc := range tc.genState.MainnetVestingAccountList {
				// check if the campaign exists for mainnet accounts
				_, ok := campaignIDMap[acc.CampaignID]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.CampaignID]; !ok {
					shares[acc.CampaignID] = types.EmptyShares()
				}
				totalShares, err := acc.GetTotalShares()
				require.NoError(t, err)

				shares[acc.CampaignID] = types.IncreaseShares(
					shares[acc.CampaignID],
					totalShares,
				)
			}

			for campaignID, share := range campaignIDMap {
				// check if the campaign shares is equal all accounts shares
				accShares, ok := shares[campaignID]
				require.True(t, ok)
				isLowerEqual := accShares.IsAllLTE(share)
				require.True(t, isLowerEqual)
			}
		})
	}
}
