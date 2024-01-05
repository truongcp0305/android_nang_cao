package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type keyData struct {
	key         string
	expriedTime time.Time
}

var aeskey = NewKey()

func NewKey() keyData {
	b := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	key := keyData{
		key: string(b),
	}
	key.expriedTime = time.Now().Add(30 * time.Minute)
	return key
}

func AesKey() string {
	return aeskey.key
}

func Encrypt(data string) (string, error) {
	if aeskey.key == "" || aeskey.expriedTime.Before(time.Now()) {
		aeskey = NewKey()
	}
	key := []byte(aeskey.key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	data = string(PKCS7Padding([]byte(data), aes.BlockSize))
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(data))
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Decrypt(encrypted string) (result string, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("Panic : %v", r)
		}
	}()
	key := []byte(aeskey.key)
	ciphertext, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Lấy iv từ dữ liệu mã hóa
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Tạo chế độ CBC
	mode := cipher.NewCBCDecrypter(block, iv)

	// Giải mã dữ liệu
	mode.CryptBlocks(ciphertext, ciphertext)
	var a = string(ciphertext)
	fmt.Print(a)

	return string(PKCS7UnPadding(ciphertext)), nil
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding xóa padding PKCS7 từ dữ liệu
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
