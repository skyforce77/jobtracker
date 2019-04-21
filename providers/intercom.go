package providers

type intercom struct {
	greenhouse
}

// NewInterCom returns a new provider
func NewInterCom() Provider {
	return &intercom{
		greenhouse{
			"InterCom",
			"intercom",
		},
	}
}
