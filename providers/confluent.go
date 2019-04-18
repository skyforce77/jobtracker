package providers

type confluent struct {
	lever
}

func NewConfluent() *confluent {
	return &confluent{
		lever{
			"Confluent",
			"confluent",
		},
	}
}
