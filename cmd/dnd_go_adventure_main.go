package main

import (
	"fmt"

	dice "github.com/conorkenn/dnd_go_adventure/pkg"
)

func main() {
	dice := dice.D100

	result := dice.Roll()
	fmt.Print(result)
}
