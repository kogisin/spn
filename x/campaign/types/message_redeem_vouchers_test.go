package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgRedeemVouchers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgRedeemVouchers{
				Creator:    "invalid_address",
				Account:    sample.AccAddress(),
				Vouchers:   sample.Coins(),
				CampaignID: 0,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid account address",
			msg: types.MsgRedeemVouchers{
				Creator:    sample.AccAddress(),
				Account:    "invalid_address",
				Vouchers:   sample.Coins(),
				CampaignID: 0,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid vouchers",
			msg: types.MsgRedeemVouchers{
				Creator:    sample.AccAddress(),
				Account:    sample.AccAddress(),
				Vouchers:   invalidCoins,
				CampaignID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "empty vouchers",
			msg: types.MsgRedeemVouchers{
				Creator:    sample.AccAddress(),
				Account:    sample.AccAddress(),
				Vouchers:   sdk.Coins{},
				CampaignID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "valid message",
			msg: types.MsgRedeemVouchers{
				Creator:    sample.AccAddress(),
				Account:    sample.AccAddress(),
				Vouchers:   sample.Coins(),
				CampaignID: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}