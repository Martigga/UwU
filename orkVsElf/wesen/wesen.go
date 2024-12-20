/* Name: Herker
 * Datum: 13.12.2024
 * Zweck: Klasse "Wesen" als Oberklasse für z. B. Ork und Elf
 */

package wesen

type Wesen struct {
	schaden int		// Schaden, der verursacht wird
	gesundheit int	// Lebenspunkte
	name string		// Name des Wesens
	abwehr int		// Abwehrkraft des Wesens
	initiative int	// Geschwindigkeit beim Angriff (1 schnell, 5 langsam)
}

/////////////// Getter und Setter //////////////////
func (w *Wesen) GetSchaden() int {
	return w.schaden
}
func (w *Wesen) SetSchaden(sp int) {
	w.schaden = sp
}

func (w *Wesen) GetGesundheit() int {
	return w.gesundheit
}
func (w *Wesen) SetGesundheit(gp int) {
	w.gesundheit = gp
}

func (w *Wesen) GetNamen() string {
	return w.name
}
func (w *Wesen) SetNamen(n string) {
	w.name = n
}

func (w *Wesen) GetAbwehr() int {
	return w.abwehr
}
func (w *Wesen) SetAbwehr(ap int) {
	w.abwehr = ap
}

func (w *Wesen) GetInitiative() int {
	return w.abwehr
}

func (w *Wesen) SetInitiative(i int) {
	w.initiative = i
}

// Reihenfolge: Schaden,Gesundheit,Abwehr,Initiative,Name
func (w *Wesen) GetAll()(int,int,int,int,string) {
	return w.schaden,w.gesundheit,w.abwehr,w.initiative,w.name
}

// Reihenfolge: Schaden,Gesundheit,Abwehr,Initiative,Name
func (w *Wesen) SetAll(s,g,a,i int, n string) {
	w.schaden = s
	w.gesundheit = g
	w.name = n
	w.abwehr = a
	w.initiative = i
}

/////////////////// Konstruktor //////////////////////////

func NeuesWesen() *Wesen {
	var w *Wesen = new(Wesen)
	return w
}


