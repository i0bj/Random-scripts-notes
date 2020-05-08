package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	fmt.Println("[+] Enter the file to encrypt: ")
	var file string
	fmt.Scanln(&file)

	// Reads file to b as []bytes
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	// The Key has to be at least 32 bytes
	key := []byte("anything")

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		fmt.Println(err)
	}

	//Writes encrypted data to file
	err = ioutil.WriteFile("Encrypted.txt", gcm.Seal(nonce, nonce, b, nil), 777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AES file created!")
}
