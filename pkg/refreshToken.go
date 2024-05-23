package pkg

import (
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/sha3"
)

var (
	symbols = []byte("1234567890qwertyuiopasdfghjklzxcvbnm!@#$^&*()_+QWERTYUIOPASDFGHJKLZXCVBNM")
)

func GenerateRefreshToken(length int) string {

	h := sha3.New256()

	result := string(time.Now().UnixMilli())

	for i := 0; i < length; i++ {
		result += string(symbols[rand.Intn(len(symbols))])
	}

	h.Write([]byte(result))

	return hex.EncodeToString(h.Sum(nil))

}
