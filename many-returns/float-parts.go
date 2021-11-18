package many_returns

import "math"

func FloatParts(number float64) (integerPart int, fractionalPart float64)  {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}
