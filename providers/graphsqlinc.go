package providers

type graphsqlinc struct {
	greenhouse
}

// NewGraphSQL returns a new provider
func NewGraphSQL() Provider {
	return &graphsqlinc{
		greenhouse{
			"GraphSQL",
			"graphsqlinc",
		},
	}
}
