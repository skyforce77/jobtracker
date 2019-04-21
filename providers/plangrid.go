package providers

type plangrid struct {
	greenhouse
}

// NewPlanGrid returns a new provider
func NewPlanGrid() Provider {
	return &plangrid{
		greenhouse{
			"PlanGrid",
			"plangrid",
		},
	}
}
