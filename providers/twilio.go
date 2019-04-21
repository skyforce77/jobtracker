package providers

type twilio struct {
	greenhouse
}

// NewTwilio returns a new provider
func NewTwilio() Provider {
	return &twilio{
		greenhouse{
			"Twilio",
			"twilio",
		},
	}
}
