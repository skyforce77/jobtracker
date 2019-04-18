package providers

type doctrine struct {
	lever
}

func NewDoctrine() *doctrine {
	return &doctrine{
		lever{
			"Doctrine",
			"doctrine",
		},
	}
}
