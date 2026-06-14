package util

import (
	"fmt"
	"time"
)

func GenerateAccountNumber() string {
	return fmt.Sprintf("SEBS%d", time.Now().UnixNano())
}
