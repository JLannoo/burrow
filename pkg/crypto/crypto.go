package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"

	"github.com/jlannoo/burrow/pkg/files"
)

var salt = []byte("saltysalt")

func hashPassword(h hash.Hash, password string) []byte {
	h.Write(salt)
	h.Write([]byte(password))
	return h.Sum(nil)
}

func ComparePasswords(h hash.Hash, password string, hash []byte) bool {
	return bytes.Equal(hashPassword(h, password), hash)
}

func HashSHA256(password string) []byte {
	return hashPassword(sha256.New(), password)
}

func CompareSHA256(password string, hash []byte) bool {
	return ComparePasswords(sha256.New(), password, hash)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomKey() ([]byte, error) {
	return GenerateRandomBytes(32)
}

func GenerateUnlockKey(hashedMasterPassword string) ([]byte, error) {
	if(hashedMasterPassword == "") {
		return nil, errors.New("hashedMasterPassword is empty")
	}

	keyBytes, err := files.Manager.ReadFromSecretKeyFile()
	if err != nil {
		return nil, errors.New("your secret key file is missing, please run burrow init")
	}

	return HashSHA256(string(keyBytes) + hashedMasterPassword), nil
}

func pad(data []byte, size int) []byte {
	padSize := size - len(data)%size
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, pad...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	if length%blockSize != 0 {
		return nil, errors.New("invalid padding on input")
	}
	paddingLen := int(data[length-1])
	if paddingLen > blockSize || paddingLen == 0 {
		return nil, errors.New("invalid padding size")
	}
	for i := 0; i < paddingLen; i++ {
		if data[length-1-i] != byte(paddingLen) {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:length-paddingLen], nil
}

func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = pad(data, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	fmt.Printf("data: %v\nkey: %v\n\n", data, key)

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	fmt.Printf("data: %v\nkey: %v\niv: %v\n\n", data, key, iv)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	fmt.Printf("data: %v\nkey: %v\niv: %v\n\n", data, key, iv)

	data, err = pkcs7Unpad(data, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return data, nil
}
