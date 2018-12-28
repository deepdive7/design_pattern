package behavioral

import "testing"

// template method design pattern
func TestTemplate(t *testing.T) {
	var p *Person = &Person{}

	p.Specific = &Boy{}
	p.SetName("qibin")
	p.Out()

	p.Specific = &Girl{}
	p.SetName("loader")
	p.Out()
}
