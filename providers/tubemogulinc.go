package providers

type tubemogulinc struct {
	greenhouse
}

// NewTubeMogul returns a new provider
func NewTubeMogul() Provider {
	return &tubemogulinc{
		greenhouse{
			"TubeMogul",
			"tubemogulinc",
		},
	}
}
