package providers

type chartboost struct {
	greenhouse
}

// NewChartboost returns a new provider
func NewChartboost() Provider {
	return &chartboost{
		greenhouse{
			"Chartboost",
			"chartboost",
		},
	}
}
