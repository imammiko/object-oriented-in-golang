package agredrive

type Transmission struct {
	tType string
}

func (t *Transmission) GetType() string {
	return t.tType
}

func (t *Transmission) SetType(tType string) {
	t.tType = tType
}
