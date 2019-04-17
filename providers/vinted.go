package providers

type vinted struct {
	lever
}

func NewVinted() *vinted {
	return &vinted{
		lever{
			"Vinted",
			"vinted",
		},
	}
}