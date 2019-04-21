package providers

type hellosign struct {
	greenhouse
}

// NewHelloSign returns a new provider
func NewHelloSign() Provider {
	return &hellosign{
		greenhouse{
			"HelloSign",
			"hellosign",
		},
	}
}
