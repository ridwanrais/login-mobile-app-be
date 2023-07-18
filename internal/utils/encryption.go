package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"strings"
)

func EncryptString(key []byte, plaintext string) (string, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Generate a random Initialization Vector (IV)
	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Pad the plaintext to a multiple of the block size
	paddedPlaintext := padPlaintext(plaintext, aes.BlockSize)

	// Create a new cipher using the AES cipher block in CBC mode
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(paddedPlaintext))

	// Prepend the IV to the ciphertext
	copy(ciphertext[:aes.BlockSize], iv)

	// Encode the ciphertext as base64 before returning it
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Helper function to pad the plaintext
func padPlaintext(plaintext string, blockSize int) string {
	padding := blockSize - len(plaintext)%blockSize
	padText := string(padding)
	return plaintext + strings.Repeat(padText, padding)
}

func DecryptString(key []byte, ciphertext string) (string, error) {
	// Decode the base64-encoded ciphertext
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Separate the IV and the actual ciphertext
	iv := decodedCiphertext[:aes.BlockSize]
	ciphertextBytes := decodedCiphertext[aes.BlockSize:]

	// Create a new cipher using the AES cipher block in CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	// Remove the padding from the plaintext
	plaintext := removePadding(string(ciphertextBytes))

	return plaintext, nil
}

// Helper function to remove the padding from the plaintext
func removePadding(plaintext string) string {
	lastIndex := len(plaintext) - 1
	padding := int(plaintext[lastIndex])
	return plaintext[:lastIndex-padding+1]
}
