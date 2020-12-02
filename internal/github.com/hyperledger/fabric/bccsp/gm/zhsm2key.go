package gm

import (
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	"github.com/Hyperledger-TWGC/tjfoc-gm/x509"
	"github.com/tw-bc-group/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp"
	sm2ZH "github.com/tw-bc-group/zhonghuan-ce/sm2"
)

type zhcesm2KeyAdapter struct {
	adapter *sm2ZH.KeyAdapter
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *zhcesm2KeyAdapter) Bytes() (raw []byte, err error) {
	return nil, nil
}

// SKI returns the subject key identifier of this key.
func (k *zhcesm2KeyAdapter) SKI() (ski []byte) {
	if k.adapter.KeyID() == "" {
		return nil
	}
	return []byte(k.adapter.KeyID())
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *zhcesm2KeyAdapter) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *zhcesm2KeyAdapter) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *zhcesm2KeyAdapter) PublicKey() (bccsp.Key, error) {
	return &zhcesm2PublicKey{k.adapter.PublicKey()}, nil
}

type zhcesm2PublicKey struct {
	pubKey *sm2.PublicKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *zhcesm2PublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalSm2PublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *zhcesm2PublicKey) SKI() (ski []byte) {
	if k.pubKey == nil {
		return nil
	}

	// Marshall the public key
	raw := elliptic.Marshal(k.pubKey.Curve, k.pubKey.X, k.pubKey.Y)

	// Hash it
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *zhcesm2PublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *zhcesm2PublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *zhcesm2PublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}
