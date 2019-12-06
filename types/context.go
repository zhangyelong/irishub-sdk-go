package types

import "irishub-sdk-go/nets"

type SdkContext struct {
	LcdAddr      string
	RpcAddr      string
	HttpTimeout  int // timeout in seconds
	HttpProvider *nets.HttpProvider
}

// NewSdkContext initialize a sdk config
func NewSdkContext() *SdkContext {
	return &SdkContext{}
}

// GetHttpURL get the full api url
func (config *SdkContext) GetHttpURL(path string) string {
	if len(config.LcdAddr) == 0 {
		panic("Please configure the lcd address")
	}

	return config.LcdAddr + path
}
