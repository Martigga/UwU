/* 	Name: Martin Schmal
Datum: 27.09.2024
Zweck: Verdeutlichung der Funktionsweise von Structs als
Konstruktionen von Klassen
*/

package klassen

import (
	"fmt"
)

//////////////////////// Klasse Teletubby ////////////////////////

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
// rand.Intn(100)
func zufall(max int) int {

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int(), "Unixnano von Time.Now: ", time.Now().UnixNano())
	return rand.Int()
}*/

//////////////////// Kampf ////////////////////

// func TeletubbyVsTeletubby(TeleMaxim, Teleziggi *Teletubby) {

// 	for

// }

const bSchaden = 10
const bGesundheit = 100
const bAbwehr = 2
const bInitiative = 1

func falscheBelegung(Teletubby Teletubby) {

	Teletubby.schaden = 20
	Teletubby.gesundheit = 100
	Teletubby.name = "TeleMaxim"
	Teletubby.abwehr = 2
	Teletubby.initiative = 4

	fmt.Println("Teleziggi in Funktion falschbelegung", Teletubby)
	// Teletubby wurde global geändert

}

func richtigeBelegung(Teletubby *Teletubby) {

	Teletubby.schaden = bSchaden
	Teletubby.gesundheit = bGesundheit
	Teletubby.name = "TeleMaxim"
	Teletubby.abwehr = bAbwehr
	Teletubby.initiative = bInitiative

	fmt.Println("Teleziggi in Funktion richtigebelegung", Teletubby)
	// Teletubby wurde global geändert

}

/////////// Standartverfahren zur Erstellung und Änderung von Objekten ///////////

// Objekt wird über eine Funktion erstellt

// Getter und Setter

func (Teletubby *Teletubby) GetName() string {
	return Teletubby.name
}

func (Teletubby *Teletubby) SetName(na string) {
	Teletubby.name = na
}

func (Teletubby *Teletubby) GetSchaden() int {
	return Teletubby.schaden
}

func (Teletubby *Teletubby) SetSchaden(sp int) {
	Teletubby.schaden = sp
}

func (Teletubby *Teletubby) GetGesundheit() int {
	return Teletubby.gesundheit
}

func (Teletubby *Teletubby) SetGesundheit(hp int) {
	Teletubby.gesundheit = hp
}

func (Teletubby *Teletubby) GetAbwehr() int {
	return Teletubby.abwehr
}

func (Teletubby *Teletubby) SetAbwehr(def int) {
	Teletubby.abwehr = def
}

func (Teletubby *Teletubby) GetInitiative() int {
	return Teletubby.initiative
}

func (Teletubby *Teletubby) SetInitiative(init int) {
	Teletubby.initiative = init
}

func NeuerTeletubby() *Teletubby {
	var teletubby *Teletubby = new(Teletubby)
	return teletubby
}

func TeletubbyFreiErstellen() *Teletubby {
	var Teletubby = NeuerTeletubby()

	var wert int
	var symb string

	fmt.Print("Lege deinen Namen fest: ")
	fmt.Scanln(&symb)
	Teletubby.SetName(symb)

	fmt.Print("Lege deinen Schaden fest: ")
	fmt.Scanln(&wert)
	Teletubby.SetSchaden(wert)

	fmt.Print("Lege deine Gesundheit fest: ")
	fmt.Scanln(&wert)
	Teletubby.SetGesundheit(wert)

	fmt.Print("Lege deine Abwehr fest: ")
	fmt.Scanln(&wert)
	Teletubby.SetAbwehr(wert)

	fmt.Print("Lege deine Initiative fest: ")
	fmt.Scanln(&wert)
	Teletubby.SetInitiative(wert)

	return Teletubby
}

/*
func main() {

	// Erstellung TeleMaxim als objekt der Klasse Teletubby
	var TeleMaxim Teletubby
	fmt.Println("Der erste Teletubby TeleMaxim", TeleMaxim)

	falscheBelegung(TeleMaxim)
	fmt.Println("TeleMaxim in der main", TeleMaxim) // TeleMaxim wurde nicht synchronisiert

	// Erstellung von Teleziggi NUR als Adressvariable
	var Teleziggi *Teletubby

	fmt.Println("\nTeleziggi nach der Erstellung", Teleziggi)

	// richtigeBelegung(Teleziggi)	// Teleziggi enthält noch keine Adresse, Funktion dubktioniert nicht

	Teleziggi = new(Teletubby)
	fmt.Println("\nTeleziggi nach der Initialisierung", Teleziggi)
	richtigeBelegung(Teleziggi)
	fmt.Println("Teleziggi in der main", Teleziggi) // Teleziggi wurde lokal geändert
	fmt.Println("Adresse von Teleziggi", &Teleziggi)
	fmt.Println("Inhalt an der Adresse von Teleziggi", *Teleziggi)
	fmt.Println("Inhalt de Adressevariablen Teleziggi", Teleziggi) //

	//Teletubby mit Funktion erstellen
	var Telebussy = NeuerTeletubby()
	fmt.Println("\nTeletubby")
	fmt.Println(Telebussy)
	Telebussy.SetSchaden(10000)
	Telebussy.GetSchaden()
	fmt.Println("Schaden: ", Telebussy.GetSchaden())

	//Teletubby frei nach Nutzer erstellen

	var Telemumu = TeletubbyFreiErstellen()
	fmt.Println(Telemumu)

}
*/
