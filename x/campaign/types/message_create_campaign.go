package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCampaign{}

func NewMsgCreateCampaign(
	coordinator string,
	campaignName string,
	totalSupply sdk.Coins,
	dynamicShares bool,
) *MsgCreateCampaign {
	return &MsgCreateCampaign{
		Coordinator:   coordinator,
		CampaignName:  campaignName,
		TotalSupply:   totalSupply,
		DynamicShares: dynamicShares,
	}
}

func (msg *MsgCreateCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateCampaign) Type() string {
	return "CreateCampaign"
}

func (msg *MsgCreateCampaign) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgCreateCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if err := CheckCampaignName(msg.CampaignName); err != nil {
		return sdkerrors.Wrap(ErrInvalidCampaignName, err.Error())
	}

	if !msg.TotalSupply.IsValid() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is not a valid Coins object")
	}

	if msg.TotalSupply.Empty() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is empty")
	}

	return nil
}