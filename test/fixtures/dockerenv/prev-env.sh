#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This file contains environment overrides to enable testing
# against the latest pre-release target.

export FABRIC_FIXTURE_VERSION="v1.4"
export FABRIC_CRYPTOCONFIG_VERSION="v1"

export FABRIC_CA_FIXTURE_TAG="latest"
export FABRIC_ORDERER_FIXTURE_TAG="latest"
export FABRIC_PEER_FIXTURE_TAG="latest"
export FABRIC_BUILDER_FIXTURE_TAG="latest"
export FABRIC_BASEOS_FIXTURE_TAG="0.4.20"

export FABRIC_CA_FIXTURE_IMAGE=hyperledger/fabric-ca-gm
export FABRIC_ORDERER_FIXTURE_IMAGE=hyperledger/fabric-orderer-gm
export FABRIC_PEER_FIXTURE_IMAGE=hyperledger/fabric-peer-gm
export FABRIC_BUILDER_FIXTURE_IMAGE=hyperledger/fabric-ccenv-gm
