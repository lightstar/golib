package context_test

type Params struct {
}

func NewParams() *Params {
	return &Params{}
}

func (params *Params) ByName(name string) string {
	return name
}
