package providers

type paloAltoNetworks struct {
	jobVite
}

func NewPaloAltoNetworks() *paloAltoNetworks {
	return &paloAltoNetworks{
		jobVite{
			"Palo Alto Networks",
			"https://jobs.jobvite.com/paloaltonetworks/jobs/all-jobs",
		},
	}
}
