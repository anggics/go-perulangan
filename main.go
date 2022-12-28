package main

import (
	"fmt"
	switchcase "go-perulangan/switchCase"
)

func main() {
	// //perulangan for 1 sampai 10
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(" func for ::", i)
	// }

	// //whlie dalam go
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("index ", i)
	// 	i++
	// }

	// // for range
	// StringTitle := "Fundamental Golang Perulangan"
	// for index, letter := range StringTitle {
	// 	fmt.Println("index ::", index, "letter :: ", string(letter))
	// }

	// result := switchcase.SwitchCase(20)
	// fmt.Println(result)

	result := switchcase.SwitchCaseFallthrough(2)
	fmt.Println(result)

}
