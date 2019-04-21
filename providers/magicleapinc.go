package providers

type magicleapinc struct {
	greenhouse
}

// NewMagicLeap returns a new provider
func NewMagicLeap() Provider {
	return &magicleapinc{
		greenhouse{
			"MagicLeap",
			"magicleapinc",
		},
	}
}
