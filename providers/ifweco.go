package providers

type ifweco struct {
	greenhouse
}

// NewIfWe returns a new provider
func NewIfWe() Provider {
	return &ifweco{
		greenhouse{
			"IfWe",
			"ifweco",
		},
	}
}
