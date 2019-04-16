package providers

type blizzard struct {
	jobVite
}

func NewBlizzard() *blizzard {
	return &blizzard{
		jobVite{
			"Blizzard",
			"https://jobs.jobvite.com/blizzard/jobs",
		},
	}
}
