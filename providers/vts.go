package providers

type vts struct {
	greenhouse
}

// NewVTS returns a new provider
func NewVTS() Provider {
	return &vts{
		greenhouse{
			"VTS",
			"vts",
		},
	}
}
