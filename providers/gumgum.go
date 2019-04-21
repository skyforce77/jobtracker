package providers

type gumgum struct {
	jobVite
}

// NewGumGum returns a new provider
func NewGumGum() Provider {
	return &gumgum{
		jobVite{
			"GumGum",
			"https://jobs.jobvite.com/gumgum/jobs",
		},
	}
}
