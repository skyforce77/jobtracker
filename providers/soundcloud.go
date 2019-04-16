package providers

type soundcloud struct {
	jobVite
}

func NewSoundcloud() *soundcloud {
	return &soundcloud{
		jobVite{
			"Soundcloud",
			"https://jobs.jobvite.com/soundcloud/jobs",
		},
	}
}
