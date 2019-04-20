package balance

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Balance interface {
	GetIndex(string, int) int
}
