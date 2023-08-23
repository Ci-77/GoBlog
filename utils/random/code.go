package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var stringCode = ""

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("#{rand.Intn(10000)}")
}
