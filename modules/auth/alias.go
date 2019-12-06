package auth

import "github.com/irisnet/irishub/app/v1/auth"

type (
	Account = auth.Account
	AccountDecoder = auth.AccountDecoder
)

var (
	RegisterCodec = auth.RegisterCodec
)