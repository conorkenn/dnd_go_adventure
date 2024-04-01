package dice_test

import (
	"testing"

	dice "github.com/conorkenn/dnd_go_adventure/pkg"
)

func TestDice(t *testing.T) {
	t.Run("d4", func(t *testing.T) {
		rollTester(t, dice.D4, 4)
	})
	t.Run("d6", func(t *testing.T) {
		rollTester(t, dice.D6, 6)
	})
	t.Run("d8", func(t *testing.T) {
		rollTester(t, dice.D8, 8)
	})
	t.Run("d10", func(t *testing.T) {
		rollTester(t, dice.D6, 10)
	})
	t.Run("d12", func(t *testing.T) {
		rollTester(t, dice.D6, 12)
	})
	t.Run("d20", func(t *testing.T) {
		rollTester(t, dice.D6, 20)
	})
	t.Run("d100", func(t *testing.T) {
		rollTester(t, dice.D6, 100)
	})
}

func rollTester(t testing.TB, d dice.Dice, max int) {
	result := d.Roll()
	if result > max {
		t.Errorf("Roll failed got %v expected less than %v", result, max)
	}
}
