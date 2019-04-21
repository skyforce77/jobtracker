package providers

type coupang struct {
	greenhouse
}

// NewCoupang returns a new provider
func NewCoupang() Provider {
	return &coupang{
		greenhouse{
			"Coupang",
			"coupang",
		},
	}
}
