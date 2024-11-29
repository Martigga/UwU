package waffe

type Waffe struct {
	wschaden     int
	wfarbe       string
	whaltbarkeit int
	wname        string
	wtyp         string
	winimod      int
}

// Konstruktor

func NeueWaffe() *Waffe {

	var w *Waffe = new(Waffe)
	w.wschaden = 0
	w.wfarbe = "Hautfarbe"
	w.whaltbarkeit = 1000
	w.wname = "name"
	w.wtyp = "Hände"
	w.winimod = 0

	return w

}

//////////////////// Setter ////////////////////////

func (w *Waffe) SetSchaden(x int) {
	w.wschaden = x
}

func (w *Waffe) SeFarbe(x string) {
	w.wfarbe = x
}

func (w *Waffe) SetHaltbarkeit(x int) {
	w.whaltbarkeit = x
}

func (w *Waffe) SetName(x string) {
	w.wname = x
}

func (w *Waffe) SetTyp(x string) {
	w.wtyp = x
}

func (w *Waffe) SetInitmod(x int) {
	w.winimod = x
}

func (w *Waffe) SetAll(s, h, i int, f, n, t string) {

	w.wschaden = s
	w.wfarbe = f
	w.whaltbarkeit = h
	w.wname = n
	w.wtyp = t
	w.winimod = i

}

//////////////////// Getter /////////////////////////

func (w *Waffe) GetSchaden(x int) {
	w.wschaden = x
}

func (w *Waffe) GetFarbe(x string) {
	w.wfarbe = x
}

func (w *Waffe) GetHaltbarkeit(x int) {
	w.whaltbarkeit = x
}

func (w *Waffe) GetName(x string) {
	w.wname = x
}

func (w *Waffe) GetTyp(x string) {
	w.wtyp = x
}

func (w *Waffe) GetInitmod(x int) {
	w.winimod = x
}

func (w *Waffe) GetAll() (int, int, int, string, string, string) {
	return w.wschaden, w.whaltbarkeit, w.winimod, w.wname, w.wfarbe, w.wtyp
}
