package common

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

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

func ReadFileBlob(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFileBlob(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
