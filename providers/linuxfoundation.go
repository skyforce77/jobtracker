package providers

type linuxFoundation struct {
	lever
}

// NewLinuxFoundation returns a new provider
func NewLinuxFoundation() Provider {
	return &linuxFoundation{
		lever{
			"Linux Foundation",
			"linuxfoundation.org",
		},
	}
}
