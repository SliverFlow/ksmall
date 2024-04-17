package rpc_client

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserRpc,
)
