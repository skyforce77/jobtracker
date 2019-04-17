package providers

type strait struct {
	lever
}

func NewStrait() *strait {
	return &strait{
		lever{
			"Strait",
			"strait",
		},
	}
}