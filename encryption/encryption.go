package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(key string, message string) (string, error) {
	// https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	byteMsg := []byte(message)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	// TODO: check if the key is of valid length (16, 24, or 32 bytes)
	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key string, message string) (string, error) {
	// https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func IsEncrypted(input string) bool {
	// Check if the input string is a valid base64 encoded string and has at least aes.BlockSize bytes
	if input == "" {
		return false
	}
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return false
	}
	return len(decoded) >= aes.BlockSize
}
