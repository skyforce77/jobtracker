package providers

type curalate struct {
	greenhouse
}

// NewCuralate returns a new provider
func NewCuralate() Provider {
	return &curalate{
		greenhouse{
			"Curalate",
			"curalate",
		},
	}
}
