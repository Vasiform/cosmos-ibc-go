package v7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	connectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
)

// MigrateLocalhostConnectionEnd creates the sentinel localhost connection end to enable
// localhost ibc functionality.
func MigrateLocalhostConnectionEnd(ctx sdk.Context, connectionKeeper ConnectionKeeper) {
	localhostConnection := connectionKeeper.CreateSentinelLocalhostConnection()
	connectionKeeper.SetConnection(ctx, connectiontypes.LocalhostID, localhostConnection)
}
