package random

import (
	"math/rand"
	"time"
)

// Init rand seed
func Init() {
	rand.Seed(time.Now().UnixNano())
}
