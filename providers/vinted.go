package providers

type vinted struct {
	lever
}

// NewVinted returns a new provider
func NewVinted() Provider {
	return &vinted{
		lever{
			"Vinted",
			"vinted",
		},
	}
}
