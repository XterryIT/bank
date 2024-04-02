package util

import (
	"bytes"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomString(n int) string {
	var buf bytes.Buffer

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k-1)]
		buf.WriteByte(c)
	}

	return buf.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(randomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
