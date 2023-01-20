package main

import (
	"fmt"
	"github.com/monime-lab/gomath/rational"
	"time"
)

func main() {
	t := time.Now()
	a := rational.NewFromComponents(1, 1)
	b := rational.NewFromComponents(1, 4)
	fmt.Printf("Rat: %v. Took: %s\n", a.Sub(b), time.Since(t))
}
