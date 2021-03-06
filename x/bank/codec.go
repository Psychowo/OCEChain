package bank

import (
	"github.com/OCEChain/OCEChain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/Send", nil)
	cdc.RegisterConcrete(MsgIssue{}, "cosmos-sdk/Issue", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
