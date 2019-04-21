package providers

type betclic struct {
	jobVite
}

// NewBetclic returns a new provider
func NewBetclic() Provider {
	return &betclic{
		jobVite{
			"Betclic",
			"http://jobs.jobvite.com/betclic/jobs",
		},
	}
}
