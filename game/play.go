package game

type Play struct {
	name string
}

func NewPlay(name string) *Play{
	return &Play{name: name}
}