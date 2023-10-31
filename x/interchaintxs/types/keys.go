package types

const (
	// ModuleName defines the module name
	ModuleName = "interchaintxs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interchaintxs"
)

const (
	prefixParamsKey                 = iota + 1
	prefixLastFreeRegisterICACodeID = iota + 2
)

var ParamsKey = []byte{prefixParamsKey}
var LastFreeRegisterICACodeID = []byte{prefixLastFreeRegisterICACodeID}
