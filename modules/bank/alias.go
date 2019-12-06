package bank

import "github.com/irisnet/irishub/app/v1/bank"

type (
	MsgSend = bank.MsgSend
	Input   = bank.Input
	Output  = bank.Output
)

var (
	NewMsgSend    = bank.NewMsgSend
	NewInput      = bank.NewInput
	NewOutput     = bank.NewOutput
	RegisterCodec = bank.RegisterCodec
)
