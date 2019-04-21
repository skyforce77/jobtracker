package providers

type boxinc struct {
	greenhouse
}

// NewBox returns a new provider
func NewBox() Provider {
	return &boxinc{
		greenhouse{
			"Box",
			"boxinc",
		},
	}
}
