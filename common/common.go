package common

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cakeholeDC/on-air/encryption"
	"github.com/cakeholeDC/on-air/logger"
	"github.com/zs5460/art"
	"golang.org/x/term"
	"gopkg.in/yaml.v2"
)

var log = logger.New("common")

func PrintJSON(i interface{}) {
	// Marshal the interface to JSON with indentation to make it more readable.
	// Prints to the console.
	json, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(json))
}

func PrintYAML(i interface{}) {
	// Marshal the interface to YAML to make it more readable.
	// Prints to the console.
	yaml, _ := yaml.Marshal(i)
	fmt.Println(string(yaml))
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetTerminalSize() (int, int, error) {
	// Get the size of the terminal window.
	// Returns the width and height of the terminal.
	// If it fails, it returns a default size of 80x24.
	termWidth, termHeight, err := term.GetSize(0)
	if err != nil {
		return 80, 24, err // Default size if unable to get terminal size
	}
	return termWidth, termHeight, nil
}

func PrintArtStringIfFits(content string) {
	// Prints the content as an art string if it fits in the terminal width.
	// If it does not fit, it prints the content as raw text
	artString := art.String(content)
	artLen := len(strings.SplitN(artString, "\n", -1)[1])

	termWidth, _, _ := GetTerminalSize()
	if artLen < termWidth {
		// Print the art string if it fits in the terminal width.
		fmt.Println(artString)
	} else {
		// Print the content if the art string does not fit.
		fmt.Printf("* %s\n", content)
	}
}

func ReadFileBlob(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// check if the file is encrypted
	encryptionKey := os.Getenv("ONAIR_CONFIG_ENCRYPTION_KEY")
	if encryption.IsEncrypted(string(data)) {
		// did they give us a key?
		if encryptionKey == "" {
			// if the file is encrypted by there is no key, return an error
			log.Error("encryption key is not set in the environment variable ONAIR_CONFIG_ENCRYPTION_KEY")
			return nil, fmt.Errorf("encryption key is not set in the environment variable ONAIR_CONFIG_ENCRYPTION_KEY")
		}
		// If the file is encrypted, decrypt it using the encryption key from the environment variable
		decryptedData, err := encryption.Decrypt(encryptionKey, string(data))
		if err != nil {
			return nil, err
		}
		data = []byte(decryptedData)
		// return []byte(decryptedData), nil
	} else {
		// if the file is NOT encrypted but a key is provided, return a message
		if encryptionKey != "" {
			// TODO: log or return?
			log.Warn("ONAIR_CONFIG_ENCRYPTION_KEY was provided but the file is not encrypted")
			return nil, fmt.Errorf("ONAIR_CONFIG_ENCRYPTION_KEY was provided but the file is not encrypted")
		}
	}
	// if the file is not encrypted, return the data as is
	return data, nil
}

func WriteFileBlob(filePath string, data []byte) error {
	// if there's a key, encrypt the data
	encryptionKey := os.Getenv("ONAIR_CONFIG_ENCRYPTION_KEY")
	if encryptionKey != "" {
		// Encrypt the data using the encryption key from the environment variable
		encryptedData, err := encryption.Encrypt(encryptionKey, string(data))
		if err != nil {
			fmt.Println("Error encrypting data:", err)
			return fmt.Errorf("WriteFileBlob: %w", err)
		}
		data = []byte(encryptedData)
	}
	// Write the data to the file
	return os.WriteFile(filePath, data, 0644)
}
