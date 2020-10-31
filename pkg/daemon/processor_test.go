package daemon_test

type processor struct {
	processCalled int
}

func (p *processor) Process() {
	p.processCalled++
}
