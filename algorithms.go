package main

import (
	"math"
)

func deleteNum(num, del int) int {
	i := 0
	k := 0
	s := 0.0
	for num > 0 {
		k = num%10
		num /= 10
		if k != del {
			s+=float64(k)*math.Pow(float64(10), float64(i))
			i++
		}
	}
	return int(s)
}