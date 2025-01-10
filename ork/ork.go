/* Name: Herker
 * Datum: 21.11.2024, zuletzt aktualisiert: 13.12.2024
 * Zweck: Klasse Ork
 */

package ork

import (
	"math/rand/v2"
	"start/waffe"
	"start/wesen"
)

/////////////////////// globale Konstanten und Variablen ////////////////////////////////

// Konstanten legen Werte fest, die im Code verwendet werden sollen
// hier die Basiswerte eines Orks
const bwSchaden int = 20
const bwGesundheit int = 100
const bwAbwehr int = 5
const bwInitiative int = 4

// Orknamen für eine zufällige Benennung
// (Array ist nicht als Konstante möglich)
var orknamen = [20]string{"Snaabog", "Wazgut", "Sahgigoth", "Grumgor", "Mekbag", "Duffthug", "Sniknob", "Wortrot", "Mekskab", "Podagog", "Nargog", "Dakagrot", "Usskaar", "Ugskab", "Dregbad", "Zoggog", "Nafflug", "Yambul", "Skarbad", "Grotwort"}

// Anzahl der Namen, falls sich
const anzahlnamen int = len(orknamen)

////////////////// Hilfsfunktionen ////////////////////////////

// Vor.: -
// Erg.: Eine Zufallszahl zwischen 0 und max ist zurückgegeben.
// Eff.: -
func zufall(max int) int {
	return rand.IntN(max + 1)
}

//////////////////// Klasse Ork ////////////////////////

// type erzeugt einen eigenen Datentyp
// struct erzeugt einen zusammengesetzten Datentyp
type Ork struct {
	wesen.Wesen
	owaffe *waffe.Waffe //
}

// Konstruktor
func NeuerOrk() *Ork {
	var o *Ork = new(Ork)
	o.Basiswerte()
	return o
}

/////////////////////// Methoden im package Ork ///////////////////////////

// Vor.: -
// Eff.: Das Objekt ist mit vorgegebenen Werten belegt
// Erg.: -
func (o *Ork) Basiswerte() { // (ork *Ork) bindet diese Funktion an ein Objekt vom Typ *Ork;
	// das aufrufende Objekt wird lokal als ork gespeichert
	// (d.h. in ork liegt die Adresse des Objekts)
	o.SetSchaden(bwSchaden)
	o.SetGesundheit(bwGesundheit)
	o.SetNamen(orknamen[zufall(anzahlnamen-1)]) // Zufallszahl für den Array mit Namen
	o.SetAbwehr(bwAbwehr)
	o.SetInitiative(bwInitiative)
	o.owaffe = waffe.NeueWaffe() // Die Waffe ist mit dem Konstruktor der Klasse Waffe neu geschaffen

}

///// Methoden zur Waffe

// dem Ork eine neue Waffe geben
func (o *Ork) SetWaffe(w *waffe.Waffe) {
	o.owaffe = w
}

// Ausgabe aller Werte der Waffe
func (o *Ork) GetWaffe() (int, int, int, string, string, string) {
	return o.owaffe.GetAll()
}

func (o *Ork) GetWaffenSchaden() int {
	return o.owaffe.GetSchaden()
}

func (o *Ork) GetWaffenIniMod() int {
	return o.owaffe.GetIniMod()
}

func (o *Ork) GetWaffenHaltbarkeit() int {
	return o.owaffe.GetHaltbarkeit()
}

func (o *Ork) GetWaffenFarbe() string {
	return o.owaffe.GetFarbe()
}

func (o *Ork) GetWaffenName() string {
	return o.owaffe.GetName()
}

func (o *Ork) GetWaffenTyp() string {
	return o.owaffe.GetTyp()
}
