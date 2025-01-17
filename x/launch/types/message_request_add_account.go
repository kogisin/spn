package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddAccount = "request_add_account"

var _ sdk.Msg = &MsgRequestAddAccount{}

func NewMsgRequestAddAccount(address string, launchID uint64, coins sdk.Coins) *MsgRequestAddAccount {
	return &MsgRequestAddAccount{
		Address:  address,
		LaunchID: launchID,
		Coins:    coins,
	}
}

func (msg *MsgRequestAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddAccount) Type() string {
	return TypeMsgRequestAddAccount
}

func (msg *MsgRequestAddAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if !msg.Coins.IsValid() || msg.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, msg.Address)
	}
	return nil
}
