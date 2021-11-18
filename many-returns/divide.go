package many_returns

import "errors"

func Divide(dividend float64, divisor float64) (float64, error)  {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

