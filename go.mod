// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/tw-bc-group/fabric-sdk-go-gm

go 1.14

require (
	github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20201117155542-9542a21fafd1
	github.com/Knetic/govaluate v3.0.1-0.20170926212237-aa73cfd04eeb+incompatible
	github.com/cloudflare/cfssl v1.5.0
	github.com/go-kit/kit v0.8.1-0.20190102110407-aed320776b71
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-lib-go v1.0.1-0.20181230093725-20a0acfb37ba
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.3.2
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/afero v1.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/tjfoc/gmsm v1.3.2
	github.com/tw-bc-group/fabric-gm v0.0.0-20201112022540-d3810d6681a7
	github.com/tw-bc-group/net-go-gm v0.0.0-20201030055721-5906b2d70408
	github.com/tw-bc-group/zhonghuan-ce v0.1.0
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
	golang.org/x/net v0.0.0-20201026091529-146b70c837a4
	golang.org/x/tools v0.0.0-20201023174141-c8cfbd0f21e6
	google.golang.org/grpc v1.31.0
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20201027032413-de75d571dd85 => github.com/tw-bc-group/tjfoc-gm v0.0.0-20201111115702-d6eb42f3ea58
)
