package rational

import (
	"math"
	"math/big"
)

// MaxIterations is some sane limit of iterations for precision mode
const MaxIterations = 5000

type Rational big.Rat

func (r Rational) Numerator() *big.Int {
	rat := big.Rat(r)
	return rat.Num()
}

func (r Rational) Denominator() *big.Int {
	rat := big.Rat(r)
	return rat.Denom()
}

func (r Rational) NumeratorInt64() int64 {
	return r.Numerator().Int64()
}

func (r Rational) DenominatorInt64() int64 {
	return r.Denominator().Int64()
}

func (r Rational) Add(r2 Rational) Rational {
	a := big.Rat(r)
	b := big.Rat(r2)
	res := a.Add(&a, &b)
	return Rational(*res)
}

func (r Rational) Sub(r2 Rational) Rational {
	a := big.Rat(r)
	b := big.Rat(r2)
	res := a.Sub(&a, &b)
	return Rational(*res)
}

func (r Rational) Mul(r2 Rational) Rational {
	a := big.Rat(r)
	b := big.Rat(r2)
	res := a.Mul(&a, &b)
	return Rational(*res)
}

func (r Rational) Equals(r2 Rational) bool {
	return r.Compare(r2) == 0
}

func (r Rational) Compare(r2 Rational) int {
	a, b := big.Rat(r), big.Rat(r2)
	return a.Cmp(&b)
}

func (r Rational) String() string {
	rat := big.Rat(r)
	return rat.RatString()
}

// NewWithIterations returns rational from decimal
// using `iterations` number of iterations in Continued Fraction algorithm
func NewWithIterations(val float64, iterations int64) Rational {
	return New(val, iterations, 0)
}

// NewWithPrecision returns rational from decimal
// by going as mush iterations, until next fraction is less than `stepPrecision`
func NewWithPrecision(val float64, stepPrecision float64) Rational {
	return New(val, MaxIterations, stepPrecision)
}

func NewWithDecimalPlaces(val float64, dp int) Rational {
	denominator := math.Pow10(dp)
	numerator := int64(val * denominator)
	return NewFromComponents(numerator, int64(denominator))
}

func NewFromComponents(num, den int64) Rational {
	rat := new(big.Rat)
	rat = rat.SetFrac64(num, den)
	return Rational(*rat)
}

func NewFromBigIntComponents(num, den *big.Int) Rational {
	rat := new(big.Rat)
	rat = rat.SetFrac(num, den)
	return Rational(*rat)
}

func New(val float64, iterations int64, stepPrecision float64) Rational {
	a0 := int64(math.Floor(val))
	x0 := val - float64(a0)
	rat := cf(x0, 1, iterations, stepPrecision)
	rat = rat.Add(rat, new(big.Rat).SetInt64(a0))
	return Rational(*rat)
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
