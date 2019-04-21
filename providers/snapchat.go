package providers

type snapchat struct {
	greenhouse
}

// NewSnapchat returns a new provider
func NewSnapchat() Provider {
	return &snapchat{
		greenhouse{
			"Snapchat",
			"snapchat",
		},
	}
}
