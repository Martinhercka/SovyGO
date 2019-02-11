package scrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	rnd "math/rand"
	"strings"
	"time"
)

var characters []string = strings.Split("poiuztrewqLKJHGFDSAmnbvcxPOIUZTREWQlkjhgfdsaMNBVCXY1234567890", "")

//NewPasswordHash return hash, salt
func NewPasswordHash(passwd string) (string, string) {
	salt := newSalt()
	swap := strings.Split(salt, "/")
	hash := Hasher(swap[0] + passwd + swap[1])
	return hash, salt
}

//MatchPasswordHash return bool
func MatchPasswordHash(passwd string, salt string, hash string) bool {
	swap := strings.Split(salt, "/")
	hash2 := Hasher(swap[0] + passwd + swap[1])
	if hash == hash2 {
		return true
	}
	return false
}

//Hasher simple hashing function NOT FOR RAW PASSWORDS
func Hasher(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//NewRandomPassword generate random password
func NewRandomPassword() string {
	s := rnd.NewSource(time.Now().Unix())
	r := rnd.New(s)
	var out string
	for i := 0; i < 4; i++ {
		out += characters[r.Intn(len(characters))]
	}
	out += "#"
	for i := 0; i < 4; i++ {
		out += characters[r.Intn(len(characters))]
	}
	return out
}

func newSalt() string {
	s := rnd.NewSource(time.Now().Unix())
	r := rnd.New(s)
	var out string
	for i := 0; i < 4; i++ {
		out += characters[r.Intn(len(characters))]
	}
	out += "/"
	for i := 0; i < 4; i++ {
		out += characters[r.Intn(len(characters))]
	}
	return out
}

//NewToken generate random string used as token
func NewToken() string {
	s := rnd.NewSource(time.Now().Unix())
	r := rnd.New(s)
	//tm := time.Now()
	var out string //= strconv.Itoa(tm.Hour) + strconv.Itoa(tm.Minute) + strconv.Itoa(tm.Second) + "/"
	for i := 0; i < 22; i++ {
		out += characters[r.Intn(len(characters))]
	}
	return out
}

//EncryptAES provide simple AES Encryption
func EncryptAES(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

//DecryptAES provide simple AES Decryption
func DecryptAES(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
