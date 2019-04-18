package providers

type trainline struct {
	lever
}

func NewTrainline() *trainline {
	return &trainline{
		lever{
			"Trainline",
			"trainline",
		},
	}
}
