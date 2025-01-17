package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
)

type CampaignKeeper interface {
	GetCampaign(ctx sdk.Context, id uint64) (campaigntypes.Campaign, bool)
	AddChainToCampaign(ctx sdk.Context, campaignID, launchID uint64) error
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
	GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (address string, found bool)
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
