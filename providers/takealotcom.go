package providers

type takealotcom struct {
	greenhouse
}

// NewTakealot returns a new provider
func NewTakealot() Provider {
	return &takealotcom{
		greenhouse{
			"Takealot",
			"takealotcom",
		},
	}
}
