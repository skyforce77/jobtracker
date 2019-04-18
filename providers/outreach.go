package providers

type outreach struct {
	lever
}

func NewOutreach() *outreach {
	return &outreach{
		lever{
			"Outreach",
			"outreach",
		},
	}
}
