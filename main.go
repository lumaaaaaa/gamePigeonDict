package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
	"strings"
)

// GamePigeon constant AES key and IV
const (
	AesKey = "T6wfOZgP0QgasdsgT6wfOZgP0Qgasdsg" // 32 bytes
	AesIV  = "T6wfOZgP0Qgasdsg"                 // 16 bytes
)

func main() {
	// instantiate AES cipher
	block, err := aes.NewCipher([]byte(AesKey))
	if err != nil {
		log.Fatalf("Failed to create AES cipher: %v", err)
	}

	// list files in "encrypted" directory
	files, err := os.ReadDir("encrypted")
	if err != nil {
		log.Fatalf("failed to get files in 'encrypted' directory: %v", err)
	}

	// check if there are files in the directory
	if len(files) == 0 {
		log.Fatalf("no files found in 'encrypted' directory")
	}

	os.Mkdir("decrypted", 0755)

	for _, file := range files {
		// get file name
		inputFile := file.Name()

		// read encrypted file
		encryptedData, err := os.ReadFile("encrypted/" + inputFile)
		if err != nil {
			log.Fatalf("failed to read input file: %v", err)
		}

		// ensure input size is a multiple of the block size
		if len(encryptedData)%aes.BlockSize != 0 {
			log.Fatalf("encrypted data is not a multiple of the AES block size")
		}

		// decrypt the file with AES-CBC
		mode := cipher.NewCBCDecrypter(block, []byte(AesIV))
		decrypted := make([]byte, len(encryptedData))
		mode.CryptBlocks(decrypted, encryptedData)

		// remove padding
		decrypted = decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]

		// write the decrypted data to a new file
		outputFile := inputFile[:strings.LastIndex(inputFile, ".")] + ".txt"
		outputFilePath := "decrypted/" + outputFile
		if err := os.WriteFile(outputFilePath, decrypted, 0644); err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}

		fmt.Printf("[i] decrypted '%s' to '%s', saved in 'decrypted' dir.\n", inputFile, outputFile)
	}
}
