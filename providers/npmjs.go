package providers

type npmjs struct {
	lever
}

// NewNpmjs returns a new provider
func NewNpmjs() Provider {
	return &npmjs{
		lever{
			"Npmjs",
			"npm",
		},
	}
}
