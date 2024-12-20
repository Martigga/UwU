/* Name: Herker
 * Datum: 11.10.2024, geändert: 22.11.2024
 * Zweck: Verdeutlichung der Funktionsweise von Structs als
 * 		  Konstruktionen von Klassen.
 */

package main

import (
	"fmt"
	"math/rand"

	//"time"
	"orkVsElf/ork"
)

// Vor.: -
// Erg.: Ein zufällige Zahl von 0 bis max ist zurückgegeben
// Eff.: -
func zufall(max int) int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

////////////////// Kampf //////////////////////

func OrkVsOrk(ork1, ork2 *ork.Ork) {

}

///////////////// Hilfsfunktionen /////////////////////

/*
func falscheBelegung (ork Ork) {

	ork.schaden = 20
	ork.gesundheit = 100
	ork.name = "Günther"
	ork.abwehr = 2
	ork.initiative = 4

	fmt.Println("ork1 in Funktion falscheBelegung:",ork)
	// ork wurde nur lokal geändert
}

func richtigeBelegung (ork *Ork) {
	ork.schaden = 20
	ork.gesundheit = 100
	ork.name = "Günther"
	ork.abwehr = 2
	ork.initiative = 4

	fmt.Println("ork2 in Funktion richtigeBelegung:",ork)
	// ork wurde global geändert, weil die Adresse verwendet wurde
}

*/

// Vor.: -
// Effekt.: Nutzer kann einen neuen Ork mit eigenen Werten erstellen
// Ergebnis.: Ork mit selbstgewählten Werten ist zurückgegeben
func OrkFreiErstellen() *ork.Ork {

	// notwendige Variablen
	var ork1 = ork.NeuerOrk()
	var wert int
	var text string

	fmt.Println("\n\n\n---------------------------------------------------------")
	fmt.Println("\n           Wir erstellen einen neuen Ork!")
	fmt.Println("\n---------------------------------------------------------")

	fmt.Print("\nWie soll der neue Ork heißen? : ")
	fmt.Scanln(&text)
	ork1.SetNamen(text)

	fmt.Print("Wie viel Gesundheit soll der neue Ork haben? : ")
	fmt.Scanln(&wert)
	ork1.SetGesundheit(wert)

	fmt.Print("Wie viel Abwehr soll der neue Ork haben? : ")
	fmt.Scanln(&wert)
	ork1.SetAbwehr(wert)

	fmt.Print("Wie viel Schaden soll der neue Ork verursachen? : ")
	fmt.Scanln(&wert)
	ork1.SetSchaden(wert)

	fmt.Print("Wie viel Initiative soll der neue Ork haben? : ")
	fmt.Scanln(&wert)
	ork1.SetInitiative(wert)

	return ork1
}

//////////////////////////////// MAIN //////////////////////////////////////////

func Main() {

	/*
		// Erstellung von ork1 als Objekt der Klasse Ork
		var ork1 Ork
		fmt.Println("Unser erster Ork ork1:",ork1)

		falscheBelegung(ork1)
		fmt.Println("ork1 in der main:",ork1) // ork1 wurde nicht synchronisiert


		// Erstellung von ork2 NUR als Adressvariable
		var ork2 *Ork
		fmt.Println("\nork2 nach der Erstellung:",ork2)

		// richtigeBelegung(ork2)	// ork2 enthält noch keine Adresse, Funktion funktioniert nicht

		ork2 = new(Ork)
		fmt.Println("\nork2 nach der Initialisierung:",ork2)
		richtigeBelegung(ork2)
		fmt.Println("ork2 in der main:",ork2) // ork2 wurde global geändert
											  // Durch fmt.Println() automatisch dereferenziert:
											  // Es wird der INHALT an der Adresse
											  // ausgegeben, die IN ork2 liegt
		fmt.Println("\nAdresse der Variablen ork2:",&ork2)
		fmt.Println("Inhalt an der Adresse von ork2:",*ork2)
		println("Inhalt der Adressvariablen ork2:",ork2)  // kann mit println() ausgegeben werden

	*/

	/////////// Standardverfahren zur Erstellung und Änderung von Objekten /////////////////

	// Objekt wird über eine Funktion erstellt
	var ork3 = ork.NeuerOrk()
	fmt.Println("\nork3 erstellt mit Funktion NeuerOrk:", ork3)

	// das Objekt ork3 ruft die Methode Basiswerte() auf
	ork3.Basiswerte()
	fmt.Println("ork3 nach Methode Basiswerte():", ork3)

	// Basiswerte()		// Aufruf der Funktion nicht möglich

	// Teste Getter und Setter
	fmt.Println("\nSchaden des Orks:", ork3.GetSchaden())
	fmt.Println("Schaden des Orks wird auf 25 gesetzt.")
	ork3.SetSchaden(25)
	fmt.Println("Schaden des Orks:", ork3.GetSchaden())

	// Teste die Funktion, den Ork frei zu erstellen
	// var ork4 = OrkFreiErstellen()
	// fmt.Println("\nMain: Unser neuer Ork:",ork4)

	for i := 0; i < 10; i++ {
		fmt.Println(zufall(10))
	}

	//var ork4 = OrkFreiErstellen()
	//fmt.Println("Daten des frei erstellten Orks:",ork4)

	var ork5 = ork.NeuerOrk()
	fmt.Println("Standarddaten eines Orks:", ork5)
	//fmt.Println(ork5.owaffe.GetAll()) // Waffe so nicht zugänglich
	fmt.Print("Werte der Waffe: ")
	fmt.Println(ork5.GetWaffe())

}
