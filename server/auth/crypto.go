package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Create Hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt ...
func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase))) // Create Block Cipher

	// Galois Counter Mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Cretaing Nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypted Result
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// Decrypt ...
func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))

	// Block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Galois Counter Mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Separating Noice
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypted Result
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}
