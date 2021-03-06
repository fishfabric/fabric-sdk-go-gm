package common

import (
	"crypto/x509"
	"github.com/Hyperledger-TWGC/tjfoc-gm/gmtls/gmcredentials"
	x509GM "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/client/common/verifier"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/common/providers/fab"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/core/config/comm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ConfigTLS(cert *x509.Certificate, serverName string, config fab.EndpointConfig) (grpc.DialOption, error) {
	gmtlsConfig, err := comm.GMTLSConfig(cert, serverName, config)
	if err != nil {
		tlsConfig, err := comm.TLSConfig(cert, serverName, config)
		if err != nil {
			return nil, err
		}
		//verify if certificate was expired or not yet valid
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			return verifier.VerifyPeerCertificate(rawCerts, verifiedChains)
		}
		return grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), nil
	}
	gmtlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509GM.Certificate) error {
		return verifier.VerifyPeerGMCertificate(rawCerts, verifiedChains)
	}
	return grpc.WithTransportCredentials(gmcredentials.NewTLS(gmtlsConfig)), nil
}
