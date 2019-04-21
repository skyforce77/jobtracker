package providers

type virtahealth struct {
	greenhouse
}

// NewVirtaHealth returns a new provider
func NewVirtaHealth() Provider {
	return &virtahealth{
		greenhouse{
			"VirtaHealth",
			"virtahealth",
		},
	}
}
