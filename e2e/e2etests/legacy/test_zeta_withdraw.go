package legacy

import (
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestZetaWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	// parse withdraw amount
	amount := utils.ParseBigInt(r, args[0])

	r.LegacyDepositAndApproveWZeta(amount)
	tx := r.LegacyWithdrawZeta(amount, true)

	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "zeta withdraw")
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_OutboundMined)
}
