package providers

type rosettaStone struct {
	jobVite
}

// NewRosettaStone returns a new provider
func NewRosettaStone() Provider {
	return &rosettaStone{
		jobVite{
			"Rosetta Stone",
			"https://jobs.jobvite.com/rosettastone/jobs",
		},
	}
}
