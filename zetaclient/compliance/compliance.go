// Package compliance provides functions to check for compliance of cross-chain transactions
package compliance

import (
	"github.com/rs/zerolog"

	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
	"github.com/zeta-chain/node/zetaclient/config"
)

// IsCctxRestricted returns true if the cctx involves restricted addresses
func IsCctxRestricted(cctx *crosschaintypes.CrossChainTx) bool {
	sender := cctx.InboundParams.Sender
	receiver := cctx.GetCurrentOutboundParam().Receiver

	return config.ContainRestrictedAddress(sender, receiver)
}

// PrintComplianceLog prints compliance log with fields [chain, cctx/inbound, chain, sender, receiver, token]
func PrintComplianceLog(
	inboundLogger zerolog.Logger,
	complianceLogger zerolog.Logger,
	outbound bool,
	chainID int64,
	identifier, sender, receiver, token string,
) {
	var logMsg string
	var inboundLoggerWithFields zerolog.Logger
	var complianceLoggerWithFields zerolog.Logger

	if outbound {
		// we print cctx for outbound tx
		logMsg = "Restricted address detected in cctx"
		inboundLoggerWithFields = inboundLogger.With().
			Int64("chain", chainID).
			Str("cctx", identifier).
			Str("sender", sender).
			Str("receiver", receiver).
			Str("token", token).
			Logger()
		complianceLoggerWithFields = complianceLogger.With().
			Int64("chain", chainID).
			Str("cctx", identifier).
			Str("sender", sender).
			Str("receiver", receiver).
			Str("token", token).
			Logger()
	} else {
		// we print inbound for inbound tx
		logMsg = "Restricted address detected in inbound"
		inboundLoggerWithFields = inboundLogger.With().Int64("chain", chainID).Str("inbound", identifier).Str("sender", sender).Str("receiver", receiver).Str("token", token).Logger()
		complianceLoggerWithFields = complianceLogger.With().Int64("chain", chainID).Str("inbound", identifier).Str("sender", sender).Str("receiver", receiver).Str("token", token).Logger()
	}

	inboundLoggerWithFields.Warn().Msg(logMsg)
	complianceLoggerWithFields.Warn().Msg(logMsg)
}
