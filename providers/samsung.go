package providers

type samsung struct {
	jobVite
}

func NewSamsung() *samsung {
	return &samsung{
		jobVite{
			"Samsung",
			"http://jobs.jobvite.com/samsungssi/jobs",
		},
	}
}
