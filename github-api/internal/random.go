package internal

import (
	"math/rand"
	"time"
)

func RandomBool() bool {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func RandomStatus(busy bool) string {
	if busy {
		return "online"
	}
	if RandomBool() {
		return "online"
	}
	return "offline"
}
