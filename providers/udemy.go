package providers

type udemy struct {
	lever
}

func NewUdemy() *udemy {
	return &udemy{
		lever{
			"Udemy",
			"udemy",
		},
	}
}
