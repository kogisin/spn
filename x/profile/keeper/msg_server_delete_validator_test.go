package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteValidator(t *testing.T) {
	var (
		addr1       = sample.AccAddress()
		addr2       = sample.AccAddress()
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	k.SetValidator(ctx, types.Validator{
		Address:     addr2,
		Description: &types.ValidatorDescription{},
	})
	tests := []struct {
		name string
		msg  types.MsgDeleteValidator
		err  error
	}{
		{
			name: "delete a non-existing validator",
			msg: types.MsgDeleteValidator{
				Address: addr1,
			},
			err: sdkerrors.Wrap(types.ErrValidatorNotFound,
				fmt.Sprintf("validator: %s", addr1)),
		}, {
			name: "delete an existing validator",
			msg: types.MsgDeleteValidator{
				Address: addr2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.DeleteValidator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, found := k.GetValidator(ctx, tt.msg.Address)
			require.False(t, found, "validator was not removed")
		})
	}
}