package providers

type carta struct {
	lever
}

func NewCarta() *carta {
	return &carta{
		lever{
			"Carta",
			"carta",
		},
	}
}
