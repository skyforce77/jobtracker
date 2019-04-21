package providers

type iloanbyspringleaf struct {
	greenhouse
}

// NewILoan returns a new provider
func NewILoan() Provider {
	return &iloanbyspringleaf{
		greenhouse{
			"ILoan",
			"iloanbyspringleaf",
		},
	}
}
