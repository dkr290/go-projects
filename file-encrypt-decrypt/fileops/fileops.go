package fileops

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
	"strings"
)

func EncryptFile(f string, key string) {

	// Reading plaintext file

	plainText, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("encryption read file err: %v", err.Error())
	}

	// Creating block of algorithm

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("encryption cipher err: %v", err.Error())
	}
	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("encrypption cipher GCM err: %v", err.Error())
	}
	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	// Writing ciphertext file
	err = os.WriteFile(f+".encr", cipherText, 0777)
	if err != nil {
		log.Fatalf("encryption write file err: %v", err.Error())
	}
}

func DecryptFile(f string, key string) {
	//Reading ciphertext file
	cipherText, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	// Creating block of algorithm
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("decryption cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("decryption cipher GCM err: %v", err.Error())
	}
	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	// Writing decryption content

	newFile := strings.TrimSuffix(f, ".encr")

	err = os.WriteFile(newFile, plainText, 0777)
	if err != nil {
		log.Fatalf("decryption write file err: %v", err.Error())
	}

}
