package providers

type campaignmonitor struct {
	greenhouse
}

// NewCampaignMonitor returns a new provider
func NewCampaignMonitor() Provider {
	return &campaignmonitor{
		greenhouse{
			"CampaignMonitor",
			"campaignmonitor",
		},
	}
}
