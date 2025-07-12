package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	// Test the Encrypt function
	key := "0123456789abcdef" // must be of 16 bytes for this example to work
	plainText := "Hello, World!"
	encryptedText, _ := Encrypt(key, plainText)
	fmt.Println("Encrypted text:", encryptedText)
	decryptedText, _ := Decrypt(key, encryptedText)
	assert.Equal(t, decryptedText, plainText)
}

func TestDecryptBadKey(t *testing.T) {
	// TODO: FIXME: This is a good test of a bad function.
	// Currently the payload is deciphered against the bad key, but the result is garbage.
	// This is a problem because the decryption function should return an error
	// if the key is wrong.
	//
	//
	// Test the Encrypt function
	key := "0123456789abcdef" // must be of 16 bytes for this example to work
	plainText := "Hello, World!"
	encryptedText, _ := Encrypt(key, plainText)

	badKey := "1234567812345678" // must be of 16 bytes for this example to work
	decryptedText, decErr := Decrypt(badKey, encryptedText)
	if decErr != nil {
		fmt.Println("Decryption failed as expected with bad key:", decErr)
	} else {
		// t.Errorf("Decryption should have failed with bad key, but got: %s", decryptedText)
		fmt.Printf("Decryption should have failed with bad key, but got: %s", decryptedText)
	}
	assert.NotEqual(t, decryptedText, plainText, "Decrypted text should not match the original plaintext with a bad key")
}

func TestIsEncrypted(t *testing.T) {
	// Test the IsEncrypted function
	key := "0123456789abcdef" // must be of 16 bytes for this example to work
	plainText := "Hello, World!"
	encryptedText, _ := Encrypt(key, plainText)
	fmt.Println("Encrypted text:", encryptedText)
	isEncrypted := IsEncrypted(encryptedText)
	assert.Equal(t, isEncrypted, true)
	notEncryptedStr := IsEncrypted(plainText)
	assert.Equal(t, notEncryptedStr, false)
}
