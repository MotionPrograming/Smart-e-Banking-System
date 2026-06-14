package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GenerateAccountNumber() string {

	nBig, err := rand.Int(rand.Reader, big.NewInt(9000))
	var randomNum int64 = 1234
	if err == nil {
		randomNum = nBig.Int64() + 1000
	}

	return fmt.Sprintf("SEBS%d%d", time.Now().Unix(), randomNum)
}
