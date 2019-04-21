package providers

type letv struct {
	greenhouse
}

// NewLetv returns a new provider
func NewLetv() Provider {
	return &letv{
		greenhouse{
			"Letv",
			"letv",
		},
	}
}
