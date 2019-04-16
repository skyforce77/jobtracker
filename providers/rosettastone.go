package providers

type rosettaStone struct {
	jobVite
}

func NewRosettaStone() *rosettaStone {
	return &rosettaStone{
		jobVite{
			"Rosetta Stone",
			"https://jobs.jobvite.com/rosettastone/jobs",
		},
	}
}
