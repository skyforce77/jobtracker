package providers

type warbyparker struct {
	greenhouse
}

// NewWarbyParker returns a new provider
func NewWarbyParker() Provider {
	return &warbyparker{
		greenhouse{
			"WarbyParker",
			"warbyparker",
		},
	}
}
