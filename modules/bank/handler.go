package bank

import (
	"fmt"
	"github.com/irisnet/irishub/codec"
	"irishub-sdk-go/modules/auth"
	"irishub-sdk-go/types"
)

const (
	urlGetAccount = "/bank/accounts/%s"
)

// Bank handler
type Handler struct {
	ctx *types.SdkContext
	cdc *codec.Codec
}

// NewHandler new bank handler
func NewHandler(ctx *types.SdkContext, cdc *codec.Codec) *Handler {
	RegisterCodec(cdc)
	return &Handler{
		ctx: ctx,
		cdc: cdc,
	}
}

// GetAccount get the account information on blockchain
func (handler Handler) GetAccount(address string) (auth.Account, error) {

	bz, err := handler.ctx.HttpProvider.HttpGet(handler.ctx.GetHttpURL(fmt.Sprintf(urlGetAccount, address)))
	if err != nil {
		return nil, err
	}

	var acc auth.Account
	err = handler.cdc.UnmarshalJSON(bz, &acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
