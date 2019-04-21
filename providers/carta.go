package providers

type carta struct {
	lever
}

// NewCarta returns a new provider
func NewCarta() Provider {
	return &carta{
		lever{
			"Carta",
			"carta",
		},
	}
}
