package wesen

type Wesen struct {
	bschaden    int
	bgesundheit int
	babwehr     int
	binitiative int
	bname       string
}

func Neuwesen() *Wesen {

	var b *Wesen = new(Wesen)
	b.bschaden = 20
	b.bgesundheit = 100
	b.babwehr = 2
	b.binitiative = 4
	b.bname = "Wesen"

	return b

}

//////////////////// Setter ////////////////////////

func (b *Wesen) SetSchaden(x int) {
	b.bschaden = x
}

func (b *Wesen) SetGesundheit(x int) {
	b.bgesundheit = x
}

func (b *Wesen) Setabwehr(x int) {
	b.babwehr = x
}

func (b *Wesen) SetInitiative(x int) {
	b.binitiative = x
}

func (b *Wesen) SetName(x string) {
	b.bname = x
}

func (b *Wesen) SetAll(bs, bg, ba, bi int, bn string) {

	b.bschaden = bs
	b.bgesundheit = bg
	b.babwehr = ba
	b.binitiative = bi
	b.bname = bn

}

/////////////// Getter ///////////////

func (b *Wesen) GetSchaden(x int) {
	b.bschaden = x
}

func (b *Wesen) GetGesundheit(x int) {
	b.bgesundheit = x
}

func (b *Wesen) GetAbwehr(x int) {
	b.babwehr = x
}

func (b *Wesen) GetInitiative(x int) {
	b.binitiative = x
}

func (b *Wesen) GetName(x string) {
	b.bname = x
}

func (b *Wesen) GetAll() (int, int, int, int, string) {
	return b.bschaden, b.bgesundheit, b.babwehr, b.binitiative, b.bname
}
