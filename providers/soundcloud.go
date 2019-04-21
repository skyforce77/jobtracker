package providers

type soundcloud struct {
	jobVite
}

// NewSoundcloud returns a new provider
func NewSoundcloud() Provider {
	return &soundcloud{
		jobVite{
			"Soundcloud",
			"https://jobs.jobvite.com/soundcloud/jobs",
		},
	}
}
