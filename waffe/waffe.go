/* Name: Herker
 * Datum: 13.12.2024
 * Zweck: Klasse "Waffe"
 */


package waffe

type Waffe struct {
	
	wschaden int
	wfarbe string
	whaltbarkeit int	// in Prozent
	wname string
	wtyp string
	winimod int		// Modifikator auf die Initiative
	
}

// Konstruktor mit Standardwerten
func NeueWaffe () *Waffe {
	
	var w *Waffe = new(Waffe)
	w.wschaden = 0
	w.wfarbe = "hautfarben"
	w.whaltbarkeit = 100
	w.wname = "Hände"
	w.wtyp = "Hände"
	w.winimod = 0
	
	return w
	
}

/////////////// Setter /////////////////

func (w *Waffe) SetSchaden (x int) {
	w.wschaden = x
}

func (w *Waffe) SetHaltbarkeit (x int) {
	w.whaltbarkeit = x
}

func (w *Waffe) SetIniMod (x int) {
	w.winimod = x
}

func (w *Waffe) SetFarbe (x string) {
	w.wfarbe = x
}

func (w *Waffe) SetName (x string) {
	w.wname = x
}

func (w *Waffe) SetTyp (x string) {
	w.wtyp = x
}


// Vor.: -
// Erg.: -
// Eff.: Bei der Waffe sind Schaden,Haltbarkeit,Initiativemodifikator,
//		 Farbe,Name,Typ in dieser Reihenfolge gesetzt
func (w *Waffe) SetAll (s,h,i int, f,n,t string) {
	w.wschaden = s
	w.whaltbarkeit = h
	w.winimod = i
	w.wfarbe = f
	w.wname = n
	w.wtyp = t
}



//////////////////// Getter ////////////////////

func (w *Waffe) GetSchaden () int {
	return w.wschaden
}

func (w *Waffe) GetIniMod () int {
	return w.winimod
}

func (w *Waffe) GetHaltbarkeit () int {
	return w.whaltbarkeit
}

func (w *Waffe) GetFarbe () string {
	return w.wfarbe
}

func (w *Waffe) GetName () string {
	return w.wname
}

func (w *Waffe) GetTyp () string {
	return w.wtyp
}

// Vor.: -
// Erg.: Schaden,Haltbarkeit,Initiativemodifikator,
//		 Farbe,Name,Typ sind in dieser Reihenfolge zurückgegeben
// Eff.: -
func (w *Waffe) GetAll () (int,int,int,string,string,string) {
	return w.wschaden,w.whaltbarkeit,w.winimod,w.wfarbe,w.wname,w.wtyp
}




