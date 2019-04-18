package providers

type medium struct {
	lever
}

func NewMedium() *medium {
	return &medium{
		lever{
			"Medium",
			"medium",
		},
	}
}
