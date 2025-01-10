/* 	Name: Martin Schmal
Datum: 27.09.2024
Zweck: Lokale Globale Variablen
*/

package main

import (
	"fmt"
)

var x int = 42 // globale variable = 42 (gilt für alle die keine deklariert haben)

func function1() {
	var x int
	fmt.Println("x2:", x)
}

func function2(x int) {
	fmt.Println("x3:", x)

	x++
	fmt.Println("x4:", x)
}

func function3() {
	fmt.Println("x6:", x)

	x++
	fmt.Println("x7:", x)
}

func function4() {
	fmt.Println("x9:", x)
}

func function5() int {
	var x int
	return x
}

func function6(x int) {
	fmt.Println("x11:", x)

	if x == 0 || x == 4711 {
		var x int = 11235
		fmt.Println("x12:", x)
	}

	fmt.Println("x13", x)

	for i := 0; i < 5; i++ {
		x++
	}
	fmt.Println("x14:", x)
}

func main() {

	var x int = 4711
	fmt.Println("x1:", x)

	function1()

	function2(x)
	fmt.Println("x5:", x)

	function3()

	fmt.Println("x8:", x)

	function4()

	x = function5()
	fmt.Println("x10:", x)

	function6(x)

}
