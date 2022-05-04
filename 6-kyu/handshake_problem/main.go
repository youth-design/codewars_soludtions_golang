package main

import (
	"fmt"
	"math/big"
)

func GetParticipants(handshakes int) int {
	if handshakes == 0 {
		return 1
	}
	if handshakes == 1 {
		return 2
	}
	fmt.Println(handshakes)
	for i := 3; ; i++ {
		var fact1 big.Int
		var fact2 big.Int

		fact1.MulRange(1, int64(i))
		fact2.MulRange(1, int64(i-2))

		fact1.Div(&fact1, fact2.Mul(&fact2, big.NewInt(2)))
		if fact1.Int64() >= int64(handshakes) {
			return i
		}
	}
}
