package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

// DeobfuscateStringWithKey deobfuscates a string using a key
// It XORs each byte of the obfuscated value using part of the key
// and converts it back to the original string value.
func DeobfuscateStringWithKey(obfuscatedText string, encodingKey string) (string, error) {
	encodingKeyBytes := []byte(encodingKey)
	encodingKeyLen := len(encodingKeyBytes)

	if encodingKeyLen == 0 || len(obfuscatedText) == 0 {
		return "", fmt.Errorf("encoding key or obfuscated text is empty")
	}

	obfuscatedTextBytes, err := base64.StdEncoding.DecodeString(obfuscatedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 string: %v", err)
	}

	obfuscatedTextLen := len(obfuscatedTextBytes)
	deobfuscatedTextBytes := make([]byte, obfuscatedTextLen)

	for i := 0; i < obfuscatedTextLen; i++ {
		deobfuscatedTextBytes[i] = obfuscatedTextBytes[i] ^ encodingKeyBytes[i%encodingKeyLen]
	}

	return string(deobfuscatedTextBytes), nil
}

func main() {
	secret := flag.String("secret", "", "The Secret string to deobfuscate")
	key := flag.String("key", "", "The original obfuscation Key string")
	flag.Parse()

	if *secret == "" || *key == "" {
		fmt.Println("Both '-secret <value>' and '-key <value>' flags are required")
		os.Exit(1)
	}

	deobfuscatedText, err := DeobfuscateStringWithKey(*secret, *key)
	if err != nil {
		fmt.Printf("Error deobfuscating text: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(deobfuscatedText)
}
