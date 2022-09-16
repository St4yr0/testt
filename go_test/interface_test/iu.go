package interface_test

type Eben interface {
	Eben_or_not() bool
}

type Pidarasi struct {
	Name          string
	Prepod_ebuchi bool
	Why_pidaras   string
}

type Ne_pidarasi struct {
	Name          string
	Pes_or_ne_pes bool
}

type Iu8 struct {
	Pidarasi
	Ne_pidarasi
	Kolvo_ebanatov int
}

func (Pidarasi) Eben_or_not() bool {
	return true
}
func (Ne_pidarasi) Eben_or_not() bool {
	return true
}
