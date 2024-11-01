// crypto/crypto.go
package main

import (
	"encoding/hex"
	"fmt"
)

func findPassword() {
	nonce := make([]byte, 12)
	data := "jwang34@iit.edu"

	ciphermac, _ := hex.DecodeString(
		"2d793bb434787e88d1db0f27453ac971149a6d3138591f8fa84e133805bfc748dbe9cc10d6ab7ce5b53e0b2dff6e")

	fmt.Printf("finding password for nonce=%x, data=%s, ciphermac=%x...\n",
		nonce, data, ciphermac)

	for i := 0; i < 10000; i++ {
		password := fmt.Sprintf("%04d", i)

		// Modify code below to check that if sha256(password) is a correct
		// key to decrypt the ciphertext with MAC in ciphermac, which is
		// obtained via AES-GCM with a 0 nouce and the data being jwang34@iit.edu
		//
		// Once found, show the correct password as well as the plaintext.
		if false {
			fmt.Printf("  correct password=%s\n", password)
			break
		}
	}
}

func main() {
	validateSHA256()
	validateAESGCM()
	findPassword()
}
