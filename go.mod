// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/tw-bc-group/fabric-sdk-go-gm

go 1.14

require (
	github.com/Knetic/govaluate v3.0.1-0.20170926212237-aa73cfd04eeb+incompatible
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/cloudflare/cfssl v1.4.1
	github.com/go-kit/kit v0.8.1-0.20190102110407-aed320776b71
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-lib-go v1.0.1-0.20181230093725-20a0acfb37ba
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/magiconair/properties v1.8.1 // indirect
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
	github.com/stretchr/testify v1.5.1
	github.com/tw-bc-group/fabric-ca-gm v0.0.0-20201103025508-77defe49a340
	github.com/tw-bc-group/fabric-gm v0.0.0-20201030081954-2a38b1eb3e87
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	google.golang.org/grpc v1.31.0
	gopkg.in/yaml.v2 v2.3.0
)

// Require have a local copy of fabric-gm and fabric-ca-gm under the same directory
// TODO-gm: Remove this after finishing the intigration.
replace (
	github.com/tw-bc-group/fabric-ca-gm => ../fabric-ca-gm
	github.com/tw-bc-group/fabric-gm => ../fabric-gm
)
