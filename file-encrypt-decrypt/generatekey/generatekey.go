package generatekey

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateKey() {
	key := make([]byte, 16) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}

	keyStr := hex.EncodeToString(key) //convert to string for saving
	fmt.Println(keyStr)
}
