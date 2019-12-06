package app

import (
	"github.com/irisnet/irishub/codec"
	sdk "github.com/irisnet/irishub/types"
	"irishub-sdk-go/modules/auth"
	"irishub-sdk-go/modules/bank"
	"irishub-sdk-go/nets"
	"irishub-sdk-go/types"
)

// Sdk - IRISHub SDK
type Sdk struct {
	Ctx  *types.SdkContext
	Cdc  *codec.Codec // Cdc is public since the codec is passed into the module anyways
	Bank *bank.Handler
}

// NewSdk initialize a New IRISHub SDK instance
func NewSdk() *Sdk {
	// Create the cdc with some standard codecs
	cdc := codec.New()
	auth.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Create the default context
	ctx := types.NewSdkContext()
	ctx.HttpProvider = nets.NewHttpProvider(5)

	sdk := Sdk{
		Ctx:  ctx,
		Cdc:  cdc,
		Bank: bank.NewHandler(ctx, cdc),
	}

	return &sdk
}

// WithLcdAddr set lcd address
func (sdk *Sdk) WithLcdAddr(addr string) *Sdk {
	sdk.Ctx.LcdAddr = addr
	return sdk
}

// WithRpcAddr set tendermint rpc address
func (sdk *Sdk) WithRpcAddr(addr string) *Sdk {
	sdk.Ctx.RpcAddr = addr
	return sdk
}

// WithHttpTimeout set timeout in seconds for sdk http requests
func (sdk *Sdk) WithHttpTimeout(timeout int) *Sdk {
	sdk.Ctx.HttpTimeout = timeout
	sdk.Ctx.HttpProvider = nets.NewHttpProvider(timeout)
	return sdk
}
