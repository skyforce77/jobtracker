package providers

type commercehub struct {
	greenhouse
}

// NewCommerceHub returns a new provider
func NewCommerceHub() Provider {
	return &commercehub{
		greenhouse{
			"CommerceHub",
			"commercehub",
		},
	}
}
