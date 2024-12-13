/* 	Name: Martin Schmal
Datum: 27.09.2024
Zweck: Verdeutlichung der Funktionsweise von Structs als
Konstruktionen von Klassen
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

//////////////////////// Klasse Teletubby1 ////////////////////////

// type erzeugt einen eigenen Datentyp
// struct erzeugt einen zusammengesetzen Datentyp
type Teletubby struct {
	name       string // Name vom Krieger
	schaden    int    // Verursachender Schaden
	gesundheit int    // Health Points
	abwehr     int    // Defense
	initiative int    // Geschwindigkeit beim Angriff
}

/*

const bSchaden = 10
const bGesundheit = 100
const bInitiative = 1

func richtigeBelegung(Teletubby1 *Teletubby) {

	Teletubby1.schaden = bSchaden
	Teletubby1.gesundheit = bGesundheit
	Teletubby1.name = "TeleMaxim"
	Teletubby1.initiative = bInitiative

	fmt.Println("Teleziggi in Funktion richtigebelegung", Teletubby1)
	// Teletubby1 wurde global geändert

}

func richtigeBelegung2(TeleMaxim *Teletubby) {

	TeleMaxim.schaden = bSchaden
	TeleMaxim.gesundheit = bGesundheit
	TeleMaxim.name = "TeleMaxim"
	TeleMaxim.initiative = bInitiative

	fmt.Println("Teleziggi in Funktion richtigebelegung", TeleMaxim)
	// Teletubby1 wurde global geändert

}
*/

/////////// Standartverfahren zur Erstellung und Änderung von Objekten ///////////

// Objekt wird über eine Funktion erstellt

// Getter und Setter

func (Teletubby1 *Teletubby) GetName() string {
	return Teletubby1.name
}

func (Teletubby1 *Teletubby) SetName(na string) {
	Teletubby1.name = na
}

func (Teletubby1 *Teletubby) GetSchaden() int {
	return Teletubby1.schaden
}

func (Teletubby1 *Teletubby) SetSchaden(sp int) {
	Teletubby1.schaden = sp
}

func (Teletubby1 *Teletubby) GetGesundheit() int {
	return Teletubby1.gesundheit
}

func (Teletubby1 *Teletubby) SetGesundheit(GesGetGesundheit int) {
	Teletubby1.gesundheit = GesGetGesundheit
}

func (Teletubby1 *Teletubby) GetAbwehr() int {
	return Teletubby1.abwehr
}

func (Teletubby1 *Teletubby) SetAbwehr(def int) {
	Teletubby1.abwehr = def
}

func (Teletubby1 *Teletubby) GetInitiative() int {
	return Teletubby1.initiative
}

func (Teletubby1 *Teletubby) SetInitiative(init int) {
	Teletubby1.initiative = init
}

func NeuerTeletubby() *Teletubby {
	var teletubby *Teletubby = new(Teletubby)
	return teletubby
}

func TeletubbyFreiErstellen() *Teletubby {
	var Teletubby1 = NeuerTeletubby()

	var wert int
	var symb string

	fmt.Print("Lege deinen Namen fest: ")
	fmt.Scanln(&symb)
	Teletubby1.SetName(symb)

	fmt.Print("Lege deinen Schaden fest: ")
	fmt.Scanln(&wert)
	Teletubby1.SetSchaden(wert)

	fmt.Print("Lege deine Gesundheit fest: ")
	fmt.Scanln(&wert)
	Teletubby1.SetGesundheit(wert)

	fmt.Print("Lege deine Abwehr fest: ")
	fmt.Scanln(&wert)
	Teletubby1.SetAbwehr(wert)

	fmt.Print("Lege deine Initiative fest: ")
	fmt.Scanln(&wert)
	Teletubby1.SetInitiative(wert)

	return Teletubby1
}

func TeletubbyFreiErstellen2() *Teletubby {
	var TeleMaxim = NeuerTeletubby()

	var wert int
	var symb string

	fmt.Print("Lege deinen Namen fest: ")
	fmt.Scanln(&symb)
	TeleMaxim.SetName(symb)

	fmt.Print("Lege deinen Schaden fest: ")
	fmt.Scanln(&wert)
	TeleMaxim.SetSchaden(wert)

	fmt.Print("Lege deine Gesundheit fest: ")
	fmt.Scanln(&wert)
	TeleMaxim.SetGesundheit(wert)

	fmt.Print("Lege deine Abwehr fest: ")
	fmt.Scanln(&wert)
	TeleMaxim.SetAbwehr(wert)

	fmt.Print("Lege deine Initiative fest: ")
	fmt.Scanln(&wert)
	TeleMaxim.SetInitiative(wert)

	return TeleMaxim
}

////////////// Rich Görman attack/ /////////////////

func attack(Teletubby1, TeleMaxim *Teletubby) {
	TeleMaxim.SetGesundheit(TeleMaxim.GetGesundheit() - Teletubby1.GetSchaden())
	fmt.Println(Teletubby1.GetName(), "macht", Teletubby1.GetSchaden(), "Schaden an", TeleMaxim.GetName()) ///////Funktion für den Ork
	if TeleMaxim.GetGesundheit() < 0 {
		TeleMaxim.SetGesundheit(0)
	}
}

///////////////////// Kampf ///////////////////////

func Kampf(Teletubby1, TeleMaxim *Teletubby) {
	for runden := 0; Teletubby1.GetGesundheit() > 0 && TeleMaxim.GetGesundheit() > 0; runden++ {
		var weiter bool = true

		if rand.Intn(2) == 1 {
			attack(Teletubby1, TeleMaxim)
		} else {
			attack(Teletubby1, TeleMaxim)
		}

		if Teletubby1.GetGesundheit() == 0 || TeleMaxim.GetGesundheit() == 0 {
			fmt.Println("\n Der Kampf ist vorbei!")
			weiter = false
		}

		fmt.Println(Teletubby1.GetName(), "hat", Teletubby1.GetGesundheit(), "Leben.")
		fmt.Println(TeleMaxim.GetName(), "hat", TeleMaxim.GetGesundheit(), "Leben.")

		if Teletubby1.GetGesundheit() != 0 && TeleMaxim.GetGesundheit() != 0 {
			var answer string
			fmt.Println("\nMöchtest du eine weitere Runde kämpfen?")
			fmt.Scanln(&answer)
			if strings.ToLower(answer) == "nein" {
				weiter = false
			} else {
				fmt.Println("Zu Schade, du hast nicht nein geschrieben. Es gibt eine weitere Runde.")
			}
			if !weiter {
				fmt.Println("Der Kampf ist Vorbei: ", Teletubby1.GetName(), "hat", Teletubby1.GetGesundheit(), "Leben.")
				fmt.Println("                      ", TeleMaxim.GetName(), "hat", TeleMaxim.GetGesundheit(), "Leben.")
				break
			}
		}
	}
}

func main() {

	// Erstellung TeleMaxim als objekt der Klasse Teletubby1
	var TeleMaxim *Teletubby
	fmt.Println("Der erste Teletubby TeleMaxim", TeleMaxim)

	// Erstellung von Teleziggi NUR als Adressvariable
	var Teletubby1 *Teletubby

	fmt.Println("\nTeleziggi nach der Erstellung", Teletubby1)

	// richtigeBelegung(Teleziggi)	// Teleziggi enthält noch keine Adresse, Funktion dubktioniert nicht

	/*
		Teleziggi = new(Teletubby1)
		fmt.Println("\nTeleziggi nach der Initialisierung", Teleziggi)
		fmt.Println("Teleziggi in der main", Teleziggi) // Teleziggi wurde lokal geändert
		fmt.Println("Adresse von Teleziggi", &Teleziggi)
		fmt.Println("Inhalt an der Adresse von Teleziggi", *Teleziggi)
		fmt.Println("Inhalt de Adressevariablen Teleziggi", Teleziggi) //
	*/

	//Teletubby1 mit Funktion erstellen
	var Telebussy = NeuerTeletubby()
	fmt.Println("\nTeletubby")
	fmt.Println(Telebussy)
	Telebussy.GetSchaden()
	fmt.Println("Schaden: ", Telebussy.GetSchaden())

	//Teletubby1 frei nach Nutzer erstellen

	var Telemumu = TeletubbyFreiErstellen()
	fmt.Println(Telemumu)

	var answer string
	fmt.Print("Möchtest du, dass -", Teletubby1, "- und -", TeleMaxim, "- kämpfen? Ja oder Nein?")
	fmt.Scanln(&answer)
	if strings.ToLower(answer) == "ja" {
		Kampf(Teletubby1, TeleMaxim)
	}

}
