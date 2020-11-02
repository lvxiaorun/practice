package design_mode

type People struct {
	Name string
	Like string
}

type Option func(p *People)

func WithName(name string) Option {
	return func(p *People) {
		p.Name = name
	}
}

func WithLike(like string) Option {
	return func(p *People) {
		p.Like = like
	}
}
func NewPeople() *People {
	return &People{Name: "lv", Like: "zt"}
}
