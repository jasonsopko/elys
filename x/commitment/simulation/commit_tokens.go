package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/elys-network/elys/x/commitment/keeper"
	"github.com/elys-network/elys/x/commitment/types"
)

func SimulateMsgCommitClaimedRewards(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		//simAccount, _ := simtypes.RandomAcc(r, accs)
		//msg := &types.MsgCommitClaimedRewards{
		//	Creator: simAccount.Address.String(),
		//}

		// TODO: Handling the CommitClaimedRewards simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(&types.MsgCommitClaimedRewards{}), "CommitClaimedRewards simulation not implemented"), nil, nil
	}
}
