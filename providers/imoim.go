package providers

type imoim struct {
	greenhouse
}

// NewImoDotIm returns a new provider
func NewImoDotIm() Provider {
	return &imoim{
		greenhouse{
			"ImoDotIm",
			"imoim",
		},
	}
}
