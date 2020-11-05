// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/tw-bc-group/fabric-sdk-go-gm/test/integration

replace github.com/tw-bc-group/fabric-sdk-go-gm => ../../

// Require have a local copy of fabric-gm and fabric-ca-gm under the same directory
// TODO-gm: Remove this after finishing the intigration.
replace (
	github.com/tw-bc-group/fabric-ca-gm => ../../../fabric-ca-gm
	github.com/tw-bc-group/fabric-gm => ../../../fabric-gm
)

require (
	github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20201027032413-de75d571dd85
	github.com/golang/protobuf v1.4.2
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.5.1
	github.com/tw-bc-group/fabric-ca-gm v0.0.0-00010101000000-000000000000
	github.com/tw-bc-group/fabric-gm v0.0.0-20201030081954-2a38b1eb3e87
	github.com/tw-bc-group/fabric-sdk-go-gm v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.31.0
)

go 1.14
