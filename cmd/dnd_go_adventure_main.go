package main

import (
	"fmt"

	"github.com/conorkenn/dnd_go_adventure/internal/pkg/alignment"
)

func main() {
	a, _ := alignment.GetAlignments("CE")
	fmt.Println(a.Name())
	fmt.Println(a.Description())
}
