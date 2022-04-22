package decorator

type veggeMania struct {
}

func (p *veggeMania) getPrice() int {
	return 15
}
