package providers

type shiphawk struct {
	greenhouse
}

// NewShipHawk returns a new provider
func NewShipHawk() Provider {
	return &shiphawk{
		greenhouse{
			"ShipHawk",
			"shiphawk",
		},
	}
}
