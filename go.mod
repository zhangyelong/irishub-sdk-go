module irishub-sdk-go

go 1.13

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/irisnet/irishub v0.16.0
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/prometheus/client_model v0.0.0-20191202183732-d1d2010b5bee // indirect
	github.com/prometheus/common v0.7.0 // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
)

replace (
	github.com/tendermint/iavl => github.com/irisnet/iavl v0.12.3
	github.com/tendermint/tendermint => github.com/irisnet/tendermint v0.32.1
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
)
