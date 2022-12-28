package switchcase

import (
	"strconv"
)

func SwitchCase(point int) string {
	var result string
	switch point % 2 {
	case 0:
		result = "Apakah Nilai point :" + strconv.Itoa(point) + " adalah genap :" + strconv.FormatBool(true)
	case 1:
		result = "Apakah Nilai point :" + strconv.Itoa(point) + " adalah ganjil :" + strconv.FormatBool(true)
	default:
		result = "Default"
	}
	return result
}

func SwitchCaseFallthrough(point int) string {
	var result string
	switch {
	case point < 1, point == 0:
		result = "Apakah Nilai point sama dengan 0 atau kurang dari 1 = nilai point " + strconv.Itoa(point) + " adalah :" + strconv.FormatBool(true)
	case point == 1:
		result = "Apakah Nilai point sama dengan 1 atau lebih dari 0 = nilai point " + strconv.Itoa(point) + " adalah :" + strconv.FormatBool(true)
		fallthrough
	default:
		result = "Default point " + strconv.Itoa(point)
	}
	return result
}
