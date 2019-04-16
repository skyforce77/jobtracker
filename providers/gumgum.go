package providers

type gumgum struct {
	jobVite
}

func NewGumGum() *gumgum {
	return &gumgum{
		jobVite{
			"GumGum",
			"https://jobs.jobvite.com/gumgum/jobs",
		},
	}
}
