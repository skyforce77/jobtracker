package providers

type kespryinc struct {
	greenhouse
}

// NewKespry returns a new provider
func NewKespry() Provider {
	return &kespryinc{
		greenhouse{
			"Kespry",
			"kespryinc",
		},
	}
}
