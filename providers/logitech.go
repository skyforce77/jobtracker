package providers

type logitech struct {
	jobVite
}

// NewLogitech returns a new provider
func NewLogitech() Provider {
	return &logitech{
		jobVite{
			"Logitech",
			"https://jobs.jobvite.com/logitech/jobs",
		},
	}
}
