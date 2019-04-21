package providers

type teralytics struct {
	greenhouse
}

// NewTeralytics returns a new provider
func NewTeralytics() Provider {
	return &teralytics{
		greenhouse{
			"Teralytics",
			"teralytics",
		},
	}
}
