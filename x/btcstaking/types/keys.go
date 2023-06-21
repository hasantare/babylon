package types

const (
	// ModuleName defines the module name
	ModuleName = "btcstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btcstaking"
)

var (
	BTCValidatorKey  = []byte{0x01} // key prefix for the BTC validators
	BTCDelegationKey = []byte{0x02} // key prefix for the BTC delegations
	ParamsKey        = []byte{0x03} // key prefix for the parameters
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
