package providers

type hottopic struct {
	lever
}

func NewHottopic() *hottopic {
	return &hottopic{
		lever{
			"Hot Topic",
			"hottopic",
		},
	}
}
