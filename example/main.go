package main

import (
	"fmt"
	"github.com/monime-lab/gomath/rational"
	"time"
)

func main() {
	t := time.Now()
	a := rational.NewFromComponents(1, 5)
	b := rational.NewFromComponents(20, 100)
	fmt.Printf("Rat: %v. Took: %s\n", a.Equals(b), time.Since(t))
}
