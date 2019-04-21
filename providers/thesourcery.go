package providers

type thesourcery struct {
	greenhouse
}

// NewTheSourcery returns a new provider
func NewTheSourcery() Provider {
	return &thesourcery{
		greenhouse{
			"TheSourcery",
			"thesourcery",
		},
	}
}
