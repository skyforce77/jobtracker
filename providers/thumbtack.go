package providers

type thumbtack struct {
	greenhouse
}

// NewThumbtack returns a new provider
func NewThumbtack() Provider {
	return &thumbtack{
		greenhouse{
			"Thumbtack",
			"thumbtack",
		},
	}
}
