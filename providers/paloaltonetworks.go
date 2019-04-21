package providers

type paloAltoNetworks struct {
	jobVite
}

// NewPaloAltoNetworks returns a new provider
func NewPaloAltoNetworks() Provider {
	return &paloAltoNetworks{
		jobVite{
			"Palo Alto Networks",
			"https://jobs.jobvite.com/paloaltonetworks/jobs/all-jobs",
		},
	}
}
