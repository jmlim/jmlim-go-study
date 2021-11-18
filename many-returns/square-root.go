package many_returns

import (
	"fmt"
	"math"
)

func SquareRoot(number float64) (float64, error)  {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}
