package providers

type scribd struct {
	lever
}

func NewScribd() *scribd {
	return &scribd{
		lever{
			"Scribd",
			"scribd",
		},
	}
}