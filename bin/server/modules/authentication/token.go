package authentication

import (
	"crypto"
	"crypto/cipher"
	"crypto/rsa"
	"time"

	"github.com/michalnov/SovyGo/bin/server/modules/scrypto"
)

//Token structure used as storage for authentication data
type Token struct {
	created       time.Time
	revision      time.Time
	ClientPublic  crypto.PublicKey
	serverPublic  crypto.PublicKey
	serverPrivate rsa.PrivateKey
	SymmetricKey  []byte
	desCipher     cipher.Block
	aesCipher     cipher.Block
	highSecurity  bool
	TokenBase     string
}

//NewToken create new structure of token
func NewToken() Token {
	var out Token
	out.created = time.Now()
	out.revision = out.created
	//var key rsa.PrivateKey
	err := scrypto.NewKeypair(&out.serverPrivate)
	if err != nil {
		panic(err)
	}
	key := out.serverPrivate.Public()
	out.serverPublic = key
	return out
}
