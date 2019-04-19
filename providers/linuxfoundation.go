package providers

type linuxFoundation struct {
	lever
}

func NewLinuxFoundation() *linuxFoundation {
	return &linuxFoundation{
		lever{
			"Linux Foundation",
			"linuxfoundation.org",
		},
	}
}
