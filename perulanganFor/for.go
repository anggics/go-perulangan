package perulanganfor

import (
	"fmt"
)

func GetFor(number int) int {

	for i := 1; i <= number; i++ {
		fmt.Println(" func for ::", i)
	}

	return 0
}
