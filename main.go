package main

import (
	"fmt"

	"github.com/NortPerm/SqEquation/battle"
)

func main() {
	falcon := battle.NewSpaceship(12, 7, 5, 0, 4)
	fmt.Printf("%+v", falcon)

}
