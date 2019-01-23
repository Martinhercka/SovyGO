package authentication

import (
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/michalnov/SovyGo/bin/server/modules/scrypto"
)

//Token structure used as storage for authentication data
type Token struct {
	Created         time.Time
	Revision        time.Time
	ClientPublic    rsa.PublicKey
	ServerPublic    rsa.PublicKey
	ServerPublicPem []byte
	ServerPrivate   rsa.PrivateKey
	SymmetricKey    []byte
	DesCipher       cipher.Block
	AesCipher       cipher.Block
	HighSecurity    bool
	Autheticated    bool
}

//NewToken create new structure of token
func NewToken() Token {
	var out Token
	out.Created = time.Now()
	out.Revision = out.Created
	out.Autheticated = false
	var key rsa.PublicKey
	err := scrypto.NewKeypair(&out.ServerPrivate)
	if err != nil {
		panic(err)
	}
	key = out.ServerPrivate.PublicKey
	out.ServerPublic = key
	out.ServerPublicPem = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&out.ServerPublic),
		},
	)
	return out
}
