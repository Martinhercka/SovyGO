package scrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"hash"
	"os"
)

const charset string = "poiuztrewqLKJHGFDSAmnbvcxPOIUZTREWQlkjhgfdsaMNBVCXY1234567890"

//NewKeypair genereate private key used for RSA cipher
func NewKeypair(key *rsa.PrivateKey) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	return nil
}

//RsaEncrypt Encrypt data with rsa public key
func RsaEncrypt(data []byte, label []byte, hash hash.Hash, key *rsa.PublicKey) ([]byte, error) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(hash, rng, key, data, label)
	if err != nil {
		return []byte(""), err
	}

	return ciphertext, nil
}

//RsaDecrypt Decrypt cipher to data using RSA private key
func RsaDecrypt(data []byte, label []byte, hash hash.Hash, key *rsa.PrivateKey) ([]byte, error) {
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(hash, rng, key, data, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return []byte(""), err
	}
	return plaintext, nil
}

//AesNewCipher create new aes cipher block
func AesNewCipher(key []byte) cipher.Block {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return cipher
}

//AesEncrypt not implemented
func AesEncrypt(data []byte, cipher cipher.Block) []byte {
	var out []byte
	cipher.Encrypt(out, data)
	return out
}

//AesDecrypt not implemented
func AesDecrypt(data []byte, cipher cipher.Block) []byte {
	var out []byte
	cipher.Decrypt(out, data)
	return out
}

//NewDesCipher create new Block Cipher with key from input.
func NewDesCipher(key []byte) (cipher.Block, []byte, error) {
	ede2Key := key
	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)
	out, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		//panic(err)
		return nil, []byte(""), err
	}
	return out, tripleDESKey, nil
}

/*
//NewDesEncrypt
func NewDesEncrypt(key []byte, cipherDes cipher.Block, data []byte) ([]byte, error) {

}

//NewDesDecrypt
func NewDesDecrypt(key []byte, cipherDes cipher.Block, data []byte) ([]byte, error)  {

}
*/

//DesEncrypt encrypt data into DES block cipher
func DesEncrypt(cipherDes cipher.Block, data []byte) []byte {
	var out []byte
	cipherDes.Encrypt(out, data)
	return out
}

//DesDecrypt decryption data from DES block cipher
func DesDecrypt(cipherDes cipher.Block, data []byte) []byte {
	var out []byte
	cipherDes.Decrypt(out, data)
	return out
}
