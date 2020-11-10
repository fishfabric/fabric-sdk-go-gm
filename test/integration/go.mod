// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/tw-bc-group/fabric-sdk-go-gm/test/integration

replace github.com/tw-bc-group/fabric-sdk-go-gm => ../../

replace github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20201027032413-de75d571dd85 => github.com/tw-bc-group/tjfoc-gm v0.0.0-20201109071457-3ab3cddf8830

require (
	github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20201027032413-de75d571dd85
	github.com/golang/protobuf v1.4.2
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.5.1
	//github.com/tw-bc-group/fabric-ca-gm v0.0.0-00010101000000-000000000000
	//github.com/tw-bc-group/fabric-gm v0.0.0-20201105075656-4ae0f4093603
	github.com/tw-bc-group/fabric-sdk-go-gm v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.31.0
)

go 1.14
