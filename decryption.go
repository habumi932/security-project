package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func decrypt(root string, key string) {
	// Initialize AES in GCM mode
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic("error while setting up aes")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("error while setting up gcm")
	}

	// looping through target files
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// skip if directory
		if !info.IsDir() && path[len(path)-4:] == ".enc" {
			// decrypt the file
			fmt.Println("Decrypting " + path + "...")

			// read file contents
			encrypted, err := os.ReadFile(path)
			if err == nil {
				// Decrypt bytes
				nonce := encrypted[:gcm.NonceSize()]
				encrypted = encrypted[gcm.NonceSize():]
				original, err := gcm.Open(nil, nonce, encrypted, nil)

				// write decrypted contents
				err = os.WriteFile(path[:len(path)-4], original, 0666)
				if err == nil {
					os.Remove(path) // delete the encrypted file
				} else {
					fmt.Println("error while writing contents")
				}
			} else {
				fmt.Println("error while reading file contents")
			}
		}
		return nil
	})
	fmt.Println("All files successfully decrypted.")
}
