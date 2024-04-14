package rpc_server

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserRpc,
)
