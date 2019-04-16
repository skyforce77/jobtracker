package providers

type betclic struct {
	jobVite
}

func NewBetclic() *betclic {
	return &betclic{
		jobVite{
			"Betclic",
			"http://jobs.jobvite.com/betclic/jobs",
		},
	}
}
