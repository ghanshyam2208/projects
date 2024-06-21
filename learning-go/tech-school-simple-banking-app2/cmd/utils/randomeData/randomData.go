package randomedata

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabets = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// get random int
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	str := fmt.Sprintf("%s%d", RandomString(rand.Intn(10)), RandomInt(1, 5555))
	fmt.Println(str)
	return str
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "AUS", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
