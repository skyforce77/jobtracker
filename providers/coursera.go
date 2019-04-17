package providers

type coursera struct {
	lever
}

func NewCoursera() *coursera {
	return &coursera{
		lever{
			"Coursera",
			"coursera",
		},
	}
}