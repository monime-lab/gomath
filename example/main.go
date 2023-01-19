package main

import (
	"fmt"
	"github.com/monime-lab/gomath/rational"
	"time"
)

func main() {
	t := time.Now()
	rat := rational.NewWithDecimalPlaces(0.02, 6)
	fmt.Printf("Rat: %s. Took: %s\n", rat, time.Since(t))
}
