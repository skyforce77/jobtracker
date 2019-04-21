package providers

type mongodb struct {
	greenhouse
}

// NewMongoDB returns a new provider
func NewMongoDB() Provider {
	return &mongodb{
		greenhouse{
			"MongoDB",
			"mongodb",
		},
	}
}
