# Hyperledger Fabric 国密(GM) Client SDK for Go

This SDK enables Go developers to build solutions that interact with [Hyperledger Fabric 国密(GM)](https://github.com/tw-bc-group/fabric-gm)

## Getting started

Obtain the client SDK packages for Fabric and Fabric CA.

```bash
go get github.com/hyperledger/fabric-sdk-go-gm
```

You're good to go, happy coding! Check out the examples for usage demonstrations.

### Documentation

SDK documentation can be viewed at [GoDoc](https://godoc.org/github.com/hyperledger/fabric-sdk-go-gm).

The packages intended for end developer usage are within the pkg/client folder along with the main SDK package (pkg/fabsdk).

If you wish to use the Fabric 'Gateway' programming model, then the API is in the pkg/gateway folder.

### Examples

- [Fabcar](https://github.com/tw-bc-group/fabric-samples/tree/v1.4.7-gm-notls/fabcar/go)
- [E2E Test](test/integration/e2e/end_to_end.go): Basic example that uses SDK to query and execute transaction
- [Ledger Query Test](test/integration/pkg/client/ledger/ledger_queries_test.go): Basic example that uses SDK to query a channel's underlying ledger
- [Multi Org Test](test/integration/e2e/orgs/multiple_orgs_test.go): An example that has multiple organisations involved in transaction
- [Dynamic Endorser Selection](test/integration/pkg/fabsdk/provider/sdk_provider_test.go): An example that uses dynamic endorser selection (based on chaincode policy)
<!--
- [E2E PKCS11 Test](test/integration/e2e/pkcs11/e2e_test.go): E2E Test using a PKCS11 crypto suite and configuration
- [CLI](https://github.com/securekey/fabric-examples/tree/master/fabric-cli/): An example CLI for Fabric built with the Go SDK.
- More examples needed!
-->

## Client SDK

### Current Compatibility
The SDK's integration tests run against three tagged Fabric versions:
- prev (currently v1.4.7, working on GM)
- stable (currently v2.2.0, not support GM)
- prerelease (currently disabled)

Additionally for development purposes integration tests also run against the devstable Fabric version as needed.

### Running the test suite

Obtain the client SDK packages for Fabric and Fabric CA.

```bash
git clone https://github.com/tw-bc-group/fabric-sdk-go-gm.git
```

```bash
# In the Fabric SDK Go directory
cd fabric-sdk-go-gm/

# Optional - Automatically install Go tools used by test suite
# make depend

# Running test suite
make

# Clean test suite run artifacts
make clean
```

### Go Tags
The following Go tags can be supplied to enable additional functionality:
- experimental: includes support for experimental features.

## Contributing to the Go SDK

If you want to contribute to the Go SDK, please run the test suite and submit patches for review. For general guidelines, please refer to the Fabric project's [contribution page](http://hyperledger-fabric.readthedocs.io/en/latest/CONTRIBUTING.html).

You need:

- Go 1.14
- Make
- Docker
- Docker Compose
- Git
- gobin (GO111MODULE=off go get -u github.com/myitcv/gobin)
- libtool

Notes:

- Dependencies are handled using [Go modules](https://github.com/golang/go/wiki/Modules).

### Running a portion of the test suite

```bash
# In the Fabric SDK Go directory
cd fabric-sdk-go-gm/

# Optional - Automatically install Go tools used by test suite
# make depend

# Optional - Running only code checks (linters, license, spelling, etc)
# make checks

# Running all unit tests and checks
make unit-test

# Running all integration tests
make integration-test
```

### Running package unit tests manually

```bash
# In a package directory
go test
```

### Running integration tests manually

You need:

- A working fabric and fabric-ca set up. It is recommended that you use the docker-compose file provided in `test/fixtures/dockerenv`. It is also recommended that you use the default .env settings provided in `test/fixtures/dockerenv`. See steps below.
- Customized settings in the `test/fixtures/config/config_test.yaml` in case your Hyperledger Fabric network is not running on `localhost` or is using different ports.

#### Testing with Fabric Images at Docker Hub

The test suite defaults to the latest compatible tag of fabric images at Docker Hub.
The following commands starts Fabric:

```bash
# In the Fabric SDK Go directory
cd fabric-sdk-go-gm

# Start fabric (prev tag)
make dockerenv-prev-up

# Or more generally, start fabric at a different code level (prev, stable, prerelease, devstable)
# make dockerenv-[CODELEVEL]-up
```

#### Running Integration Tests

Fabric should now be running. In a different shell, run integration tests

```bash
# In the Fabric SDK Go directory
cd fabric-sdk-go-gm

# Use script to setup parameters for integration tests and execute them
# Previously we use to have hostnames like Fabric CA server, orderer and peer pointed to localhost
# Now since we removed this now, We will be using a different configuration
make integration-tests-local

# Or more generally, run integration tests at a different code level (prev, stable, prerelease, devstable)
# and fixture target version
# FABRIC_CODELEVEL_VER=[VER] FABRIC_CODELEVEL_TAG=[CODELEVEL] make integration-tests-local
```


```bash
# Previously we use to have hostnames like Fabric CA server, orderer and peer pointed to localhost
# Now since we removed this now, We will be using a different config file config_test_local.yaml
# which has the Fabric CA server, orderer and peers pointed to localhost
# It is also possible to run integration tests using go test directly. For example:
#cd fabric-sdk-go-gm/test/integration/
#go test -args testLocal=true

#cd fabric-sdk-go-gm/test/integration/orgs
#go test -args testLocal=true 

# You should review test/scripts/integration.sh for options and details.
# Note: you should generally prefer the scripted version to setup parameters for you.
```

<!-- NA
#### Testing with Local Build of Fabric (Advanced)

Alternatively you can use a local build of Fabric using the following commands:

```bash
# Start fabric (devstable codelevel with latest docker tags)
make dockerenv-latest-up
```

## License

Hyperledger Fabric SDK Go software is licensed under the [Apache License Version 2.0](LICENSE).

---
This document is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.
-->
