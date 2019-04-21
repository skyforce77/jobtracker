package providers

type blizzard struct {
	jobVite
}

// NewBlizzard returns a new provider
func NewBlizzard() Provider {
	return &blizzard{
		jobVite{
			"Blizzard",
			"https://jobs.jobvite.com/blizzard/jobs",
		},
	}
}
