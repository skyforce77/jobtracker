package providers

type vaynermedia struct {
	greenhouse
}

// NewVaynerMedia returns a new provider
func NewVaynerMedia() Provider {
	return &vaynermedia{
		greenhouse{
			"VaynerMedia",
			"vaynermedia",
		},
	}
}
