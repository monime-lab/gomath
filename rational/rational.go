package rational

import (
	"math"
	"math/big"
)

// MaxIterations is some sane limit of iterations for precision mode
const MaxIterations = 5000

// NewWithIterations returns rational from decimal
// using `iterations` number of iterations in Continued Fraction algorithm
func NewWithIterations(val float64, iterations int64) *big.Rat {
	return New(val, iterations, 0)
}

// NewWithPrecision returns rational from decimal
// by going as mush iterations, until next fraction is less than `stepPrecision`
func NewWithPrecision(val float64, stepPrecision float64) *big.Rat {
	return New(val, MaxIterations, stepPrecision)
}

func New(val float64, iterations int64, stepPrecision float64) *big.Rat {
	a0 := int64(math.Floor(val))
	x0 := val - float64(a0)
	rat := cf(x0, 1, iterations, stepPrecision)
	return rat.Add(rat, new(big.Rat).SetInt64(a0))
}

func cf(xi float64, i int64, limit int64, stepPrecision float64) *big.Rat {
	if i >= limit || xi <= stepPrecision {
		return big.NewRat(0, 1)
	}

	inverted := 1 / xi
	aj := int64(math.Floor(inverted))
	xj := inverted - float64(aj)
	ratAJ := new(big.Rat).SetInt64(aj)
	ratNext := cf(xj, i+1, limit, stepPrecision)
	res := ratAJ.Add(ratAJ, ratNext)
	res = res.Inv(res)

	return res
}
