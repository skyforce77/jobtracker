package providers

type wistia struct {
	greenhouse
}

// NewWistia returns a new provider
func NewWistia() Provider {
	return &wistia{
		greenhouse{
			"Wistia",
			"wistia",
		},
	}
}
