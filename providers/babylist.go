package providers

type babylist struct {
	lever
}

func NewBabylist() *babylist {
	return &babylist{
		lever{
			"Babylist",
			"babylist",
		},
	}
}
