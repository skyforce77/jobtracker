package providers

type upworthy struct {
	greenhouse
}

// NewUpworthy returns a new provider
func NewUpworthy() Provider {
	return &upworthy{
		greenhouse{
			"Upworthy",
			"upworthy",
		},
	}
}
