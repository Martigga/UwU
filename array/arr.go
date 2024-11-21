package array

import "fmt"

////////////// Array //////////////

func Rawr() {

	var feld1 [8]int                         //Deklaration eines felds/Arrays
	var feld2 [5]int = [5]int{1, 2, 3, 4, 5} //Deklaration und Initialisierung mit Werten
	feld3 := [6]int{6, 5, 4, 3, 2, 1}        //schmutzige Deklaration und Initialisierung
	var feld4 [5]int
	fmt.Println(feld1)
	feld1[4] = 12
	fmt.Println(feld1[4])

	fmt.Println(feld2)
	fmt.Println(feld3)

	// feld1 = feld3 		funktioniert nicht

	for i := 0; i < len(feld3); i++ {
		feld1[i] = feld3[i]
	}

	fmt.Println(feld1)
	feld4 = feld2
	fmt.Println(feld4)
}
