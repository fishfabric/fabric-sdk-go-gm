/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package msp

import (
	"fmt"

	"github.com/cloudflare/cfssl/log"
	fabricCaUtil "github.com/tw-bc-group/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric-ca/sdkinternal/pkg/util"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/common/logging"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/common/providers/context"
	"github.com/tw-bc-group/fabric-sdk-go-gm/pkg/common/providers/msp"
)

func Example() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	username := randomUsername()

	enrollmentSecret, err := c.Register(&RegistrationRequest{Name: username})
	if err != nil {
		fmt.Printf("Register return error %s\n", err)
		return
	}

	err = c.Enroll(username, WithSecret(enrollmentSecret))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed

}

func ExampleNew() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	if c != nil {
		fmt.Println("msp client created")
	}

	// Output: msp client created
}

func ExampleWithOrg() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx, WithOrg("org1"))
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	if c != nil {
		fmt.Println("msp client created with org")
	}

	// Output: msp client created with org
}

func ExampleWithCAInstance() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx, WithCAInstance("tlsca.org1.example.com"))
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	if c != nil {
		fmt.Println("msp client created with CA Instance")
	}

	// Output: msp client created with CA Instance
}

func ExampleEnrollWithSecret() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed

}

func ExampleEnrollWithProfile() {
	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"), WithProfile("tls"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleEnrollWithType() {
	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"), WithType("x509") /*or idemix, which is not support now*/)
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleEnrollWithLabel() {
	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"), WithLabel("ForFabric"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleEnrollWithAttributeRequests() {
	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	attrs := []*AttributeRequest{{Name: "name1", Optional: true}, {Name: "name2", Optional: true}}
	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"), WithAttributeRequests(attrs))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleClient_Register() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	_, err = c.Register(&RegistrationRequest{Name: randomUsername()})
	if err != nil {
		fmt.Printf("Register return error %s\n", err)
		return
	}
	fmt.Println("register user is completed")

	// Output: register user is completed
}

func ExampleClient_Enroll() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	err = c.Enroll(randomUsername(), WithSecret("enrollmentSecret"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleClient_Reenroll() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	username := randomUsername()

	err = c.Enroll(username, WithSecret("enrollmentSecret"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}

	err = c.Reenroll(username)
	if err != nil {
		fmt.Printf("failed to reenroll user: %s\n", err)
		return
	}

	fmt.Println("reenroll user is completed")

	// Output: reenroll user is completed

}

func ExampleClient_GetSigningIdentity() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	username := randomUsername()

	err = c.Enroll(username, WithSecret("enrollmentSecret"))
	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return
	}
	enrolledUser, err := c.GetSigningIdentity(username)
	if err != nil {
		fmt.Printf("user not found %s\n", err)
		return
	}

	if enrolledUser.Identifier().ID != username {
		fmt.Println("Enrolled user name doesn't match")
		return
	}

	fmt.Println("enroll user is completed")

	// Output: enroll user is completed
}

func ExampleClient_CreateSigningIdentity() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	testPrivKey := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBG0wawIBAQQgb48r6S0fJJNZlK/L
OBZ5zd6bMVZerXMIAuSRXgfwAx6hRANCAASlZfoznuxrXH+QJYULEs/oDzpYdBXM
YauZOm9XHU6io3L5E55+4100Pp0tJb9sOTQKPj6ZE49bZzbvGoByEkiH
-----END PRIVATE KEY-----`

	testCert := `-----BEGIN CERTIFICATE-----
MIICVjCCAfugAwIBAgICEAAwCgYIKoEcz1UBg3UwTTELMAkGA1UEBhMCQ04xEjAQ
BgNVBAgMCVNoYW5nIEhhaTEUMBIGA1UECgwLZXhhbXBsZS5jb20xFDASBgNVBAMM
C2V4YW1wbGUuY29tMB4XDTIwMTEwOTA3MTIzN1oXDTIxMTEwOTA3MTIzN1owUjEL
MAkGA1UEBhMCQ04xEjAQBgNVBAgMCVNoYW5nIEhhaTEUMBIGA1UECgwLZXhhbXBs
ZS5jb20xGTAXBgNVBAMMEHRlc3QuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggq
gRzPVQGCLQNCAASlZfoznuxrXH+QJYULEs/oDzpYdBXMYauZOm9XHU6io3L5E55+
4100Pp0tJb9sOTQKPj6ZE49bZzbvGoByEkiHo4HFMIHCMAkGA1UdEwQCMAAwEQYJ
YIZIAYb4QgEBBAQDAgWgMDMGCWCGSAGG+EIBDQQmFiRPcGVuU1NMIEdlbmVyYXRl
ZCBDbGllbnQgQ2VydGlmaWNhdGUwHQYDVR0OBBYEFNFPkl0Pq2r457vSSDfs1OY1
W7aiMB8GA1UdIwQYMBaAFHMCda87dOVo2sbCmvo1Dd+uRxL5MA4GA1UdDwEB/wQE
AwIF4DAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwCgYIKoEcz1UBg3UD
SQAwRgIhAJLhxvlqLEespAZWWk/JpFbegX2pCpkim6pU5LHKGQC9AiEA6R6L+WYd
OiLBiupfD1dWkcCXr5jaff0cZouJP2qxlMQ=
-----END CERTIFICATE-----`

	// Create signing identity based on certificate and private key
	id, err := c.CreateSigningIdentity(msp.WithCert([]byte(testCert)), msp.WithPrivateKey([]byte(testPrivKey)))
	if err != nil {
		fmt.Printf("failed when creating identity based on certificate and private key: %s\n", err)
		return
	}
	if string(id.EnrollmentCertificate()) != testCert {
		fmt.Printf("certificate mismatch\n")
		return
	}

	// In this user case client might want to import keys directly into keystore
	// out of band instead of enrolling the user via SDK. User enrolment creates a cert
	// and stores it into local SDK user store, while user might not want SDK to manage certs.
	err = importPrivateKeyOutOfBand([]byte(testPrivKey), c)
	if err != nil {
		fmt.Printf("failed to import key: %s\n", err)
		return
	}

	// Create signing identity using certificate. SDK will lookup the private key based on the certificate.
	id, err = c.CreateSigningIdentity(msp.WithCert([]byte(testCert)))
	if err != nil {
		fmt.Printf("failed when creating identity using certificate: %s\n", err)
		return
	}
	if string(id.EnrollmentCertificate()) != testCert {
		fmt.Printf("certificate mismatch\n")
		return
	}

	fmt.Println("create signing identity is completed")

	// Output: create signing identity is completed
}

func importPrivateKeyOutOfBand(privateKey []byte, c *Client) error {
	_, err := fabricCaUtil.ImportBCCSPKeyFromPEMBytes([]byte(privateKey), c.ctx.CryptoSuite(), false)
	return err
}

func ExampleClient_Revoke() {

	ctx := mockClientProvider()

	// Create msp client
	c, err := New(ctx)
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	_, err = c.Revoke(&RevocationRequest{Name: "testuser"})
	if err != nil {
		fmt.Printf("revoke return error %s\n", err)
	}
	fmt.Println("revoke user is completed")

	// Output: revoke user is completed
}

func ExampleWithCA() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	results, err := c.GetAllIdentities(WithCA("CA"))
	if err != nil {
		fmt.Printf("Get identities return error %s\n", err)
		return
	}
	fmt.Printf("%d identities retrieved\n", len(results))

	// Output: 2 identities retrieved
}

func ExampleClient_CreateIdentity() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	identity, err := c.CreateIdentity(&IdentityRequest{ID: "123", Affiliation: "org2",
		Attributes: []Attribute{{Name: "attName1", Value: "attValue1"}}})
	if err != nil {
		fmt.Printf("Create identity return error %s\n", err)
		return
	}
	fmt.Printf("identity '%s' created\n", identity.ID)

	// Output: identity '123' created
}

func ExampleClient_ModifyIdentity() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	identity, err := c.ModifyIdentity(&IdentityRequest{ID: "123", Affiliation: "org2", Secret: "top-secret"})
	if err != nil {
		fmt.Printf("Modify identity return error %s\n", err)
		return
	}
	fmt.Printf("identity '%s' modified\n", identity.ID)

	// Output: identity '123' modified
}

func ExampleClient_RemoveIdentity() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	identity, err := c.RemoveIdentity(&RemoveIdentityRequest{ID: "123"})
	if err != nil {
		fmt.Printf("Remove identity return error %s\n", err)
		return
	}
	fmt.Printf("identity '%s' removed\n", identity.ID)

	// Output: identity '123' removed
}

func ExampleClient_GetIdentity() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	identity, err := c.GetIdentity("123")
	if err != nil {
		fmt.Printf("Get identity return error %s\n", err)
		return
	}
	fmt.Printf("identity '%s' retrieved\n", identity.ID)

	// Output: identity '123' retrieved
}

func ExampleClient_GetAllIdentities() {

	// Create msp client
	c, err := New(mockClientProvider())
	if err != nil {
		fmt.Println("failed to create msp client")
		return
	}

	results, err := c.GetAllIdentities()
	if err != nil {
		fmt.Printf("Get identities return error %s\n", err)
		return
	}
	fmt.Printf("%d identities retrieved\n", len(results))

	// Output: 2 identities retrieved
}

func mockClientProvider() context.ClientProvider {
	log.SetLogger(nil)
	f := testFixture{}
	sdk := f.setup()
	logging.SetLevel("fabsdk/fab", logging.ERROR)
	provider := sdk.Context()
	return provider
}
