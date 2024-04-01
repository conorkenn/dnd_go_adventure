package dice

import (
	"math/rand"
)

type Dice int

const (
	D4   Dice = 4
	D6   Dice = 6
	D8   Dice = 8
	D10  Dice = 10
	D12  Dice = 12
	D20  Dice = 20
	D100 Dice = 100
)

func (d Dice) Roll() int {
	return rand.Intn(int(d)) + 1
}
