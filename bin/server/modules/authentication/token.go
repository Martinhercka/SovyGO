package authentication

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	s "github.com/michalnov/SovyGo/bin/server/modules/structures"
)

//Token structure used as storage for authentication data
type Token struct {
	Created         time.Time
	Revision        time.Time
	ClientPublic    rsa.PublicKey
	ServerPublic    rsa.PublicKey
	ServerPublicPem []byte
	ServerPem       string
	ServerPrivate   *rsa.PrivateKey
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
	//err := scrypto.NewKeypair(&out.ServerPrivate)
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	out.ServerPrivate = priv
	if err != nil {
		panic(err)
	}
	key = out.ServerPrivate.PublicKey
	out.ServerPublic = key
	out.ServerPublicPem = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&priv.PublicKey),
		},
	)
	out.ServerPem = string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&out.ServerPrivate.PublicKey),
		},
	))
	fmt.Println(out.ServerPem)
	return out
}

//RsaDecrypt decrypt rsa encrypted data from client
func (t *Token) RsaDecrypt(env s.Envelop) error {
	rng := rand.Reader
	data, err := rsa.DecryptPKCS1v15(rng, t.ServerPrivate, env.Body)
	if err != nil {
		return err
	}
	env.Body = data
	return nil
}
