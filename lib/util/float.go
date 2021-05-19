package util

import "strconv"

func FtoS(i float64)string{
	return strconv.FormatFloat(i*100, 'f', 1,64)
}
