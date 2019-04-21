package providers

type trainline struct {
	lever
}

// NewTrainline returns a new provider
func NewTrainline() Provider {
	return &trainline{
		lever{
			"Trainline",
			"trainline",
		},
	}
}
