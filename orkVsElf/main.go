/* Name: Herker
 * Datum: 13.12.2024
 * Zweck: Main-Package des Spiels
 */

package main

import (
	"fmt"
	"math/rand/v2"
	"orkVsElf/ork"
)

// Vor.: -
// Erg.: Eine Zufallszahl zwischen 0 und max ist zurückgegeben.
// Eff.: -
func Zufall(max int) int {
	return rand.IntN(max + 1)
}

// einzelner Angriff
// Vor.: -
// Erg.: -
// Eff.: Der Angreifer a hat beim Verteidiger v die Gesundheit verringert mit
//
//	a.Schaden + a.Waffenschaden + Zufall(0-5) - v.Abwehr - Zufall(0-5)
func Angriff(a, v *ork.Ork) {

	// Hilfsvariablen
	var z1, z2 int // Zufallswerte
	var gs int     // Gesamtschaden = Schaden Ork + Schaden Waffe

	///// Angriff mit Waffe //////
	fmt.Println("\n-------------------------------------------")
	fmt.Println(a.GetNamen(), "greift an!")
	z1 = zufall(5)
	fmt.Println("Schaden", a.GetSchaden(), "+ Schaden", a.GetWaffenName(), a.GetWaffenSchaden(), "+ Glück", z1)
	fmt.Println()
	fmt.Println(v.GetNamen(), "verteidigt!")
	z2 = zufall(5)
	fmt.Println("Abwehr:", v.GetAbwehr(), "+ Glück", z2)
	// Gesamtschaden = Schaden des Orks + Waffenschaden + z1 - Abwehr - z2
	gs = a.GetSchaden() + a.GetWaffenSchaden() + z1 - v.GetAbwehr() - z2
	if gs < 0 {
		gs = 0
	} // Falls der Gesamtschaden kleiner als 0 ist, weid er auf 0 gesetzt,
	// damit der Verteidiger keine Gesundheit gewinnt

	// Gesundheit des Verteidigers wird verringert
	v.SetGesundheit(v.GetGesundheit() - gs)
	if v.GetGesundheit() < 0 {
		v.SetGesundheit(0)
	} // Verteidiger hat keine negative Gesundheit
	fmt.Println()
	if v.GetGesundheit() > 0 {
		fmt.Println(v.GetNamen(), "hat noch", v.GetGesundheit(), "Gesundheit")
	} else {
		fmt.Println("+++++", v.GetNamen(), "ist ohnmächtig geworden! +++++")
	}
}

// Vor.: -
// Erg.: -
// Eff.: Die übergebenen Orks haben rundenweise gekämpft bis einer
//
//	einen Wert gesundheit=0 hat oder der Kampf abgebrochen wurde.
func Kampf(ork1, ork2 *ork.Ork) {

	// Hilfsvariablen
	var a, v *ork.Ork     // Angreifer und Verteidiger festlegen
	var z int             // Zufallswert
	var ende bool = false // Soll der Kampf beendet werden?
	var eingabe int       // Eingabe des Benutzers

	// Hat einer der beiden schon keine Gesundheit mehr?
	if ork1.GetGesundheit() <= 0 {
		fmt.Println(ork1.GetNamen(), "kann nicht mehr Kämpfen.")
		return // Ende der Funktion
	} else if ork2.GetGesundheit() <= 0 {
		fmt.Println(ork2.GetNamen(), "kann nicht mehr Kämpfen.")
		return // Ende der Funktion
	}

	// Angreifer und Verteidiger festlegen:
	// Wer den kleineren Initiativewert hat, beginnt
	if ork1.GetInitiative() < ork2.GetInitiative() {
		a, v = ork1, ork2
	} else if ork1.GetInitiative() > ork2.GetInitiative() {
		a, v = ork2, ork1
	} else { // bei gleichem Wert, bestimmt der Zufall, wer angreift
		z = zufall(1)
		if z == 0 {
			a, v = ork1, ork2
		} else {
			a, v = ork2, ork1
		}
	} // Ende der Festlegung von a und v

	// Kampfrunden gehen so lange, bis
	// - einer der beiden Gesundheit==0 hat, oder
	// - der Kampf abgebrochen wird.
	for ende == false { // solange ende nicht auf true gesetzt wurde
		Angriff(a, v)               // erster Angriff
		if v.GetGesundheit() == 0 { // falls Verteidiger bei 0 ist
			fmt.Println("\nDer Kampf ist vorbei,", v.GetNamen(), "hat verloren!")
			fmt.Println("+++++", a.GetNamen(), "hat gewonnen! +++++")
			return // Ende der Funktion
		} else {
			Angriff(v, a)               // Gegenangriff mit vertauschen Rollen. v wird in der func Angriff zu a
			if a.GetGesundheit() == 0 { // falls Angreifer bei 0 ist
				fmt.Println("\nDer Kampf ist vorbei,", a.GetNamen(), "hat verloren!")
				fmt.Println("+++++", v.GetNamen(), "hat gewonnen! +++++")
				return // Ende der Funktion
			}
		}

		// Kampf fortführen?
		eingabe = 0                        // Der Eingabewert wird zurückgesetzt
		for eingabe != 1 && eingabe != 2 { // solange nicht 1 oder 2 eingegeben wurde (Eingabe wiederholt sich)
			fmt.Print("\nSoll der Kampf weitergehen? (1=Ja, 2=Nein) :")
			fmt.Scanln(&eingabe)
			if eingabe == 1 {
				fmt.Println("Gut, weiter geht's!")
			} else if eingabe == 2 {
				fmt.Println("\n\u2665", ork1.GetNamen(), "und", ork2.GetNamen(), "reichen sich die Hand \u2665")
				fmt.Println("\u2665 und gehen friedlich auseinander. \u2665")
				ende = true
			}
		}

	} // Ende Kampfrunden

}

///////////////////////////////////// MAIN ////////////////////////////////////

func main() {

	// Test der Zufallszahlen
	for i := 0; i < 15; i++ {
		fmt.Println(zufall(2))
	}

	// Erstellung der Orks mit Basiswerten
	var ork1 = ork.NeuerOrk()
	fmt.Println(ork1.GetAll())
	fmt.Println(ork1.GetWaffe())
	var ork2 = ork.NeuerOrk()

	// Test des Kampfes
	Kampf(ork1, ork2)

}
