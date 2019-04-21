package providers

type doordash struct {
	greenhouse
}

// NewDoorDash returns a new provider
func NewDoorDash() Provider {
	return &doordash{
		greenhouse{
			"DoorDash",
			"doordash",
		},
	}
}
