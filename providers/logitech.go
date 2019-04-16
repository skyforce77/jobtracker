package providers

type logitech struct {
	jobVite
}

func NewLogitech() *logitech {
	return &logitech{
		jobVite{
			"Logitech",
			"https://jobs.jobvite.com/logitech/jobs",
		},
	}
}
