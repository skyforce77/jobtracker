package providers

type npmjs struct {
	lever
}

func NewNpmjs() *npmjs {
	return &npmjs{
		lever{
			"Npmjs",
			"npmjs",
		},
	}
}
