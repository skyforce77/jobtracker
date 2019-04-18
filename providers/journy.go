package providers

type journy struct {
	lever
}

func NewJourny() *journy {
	return &journy{
		lever{
			"Journy",
			"gojourny",
		},
	}
}
