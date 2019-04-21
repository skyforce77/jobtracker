package providers

type eventbrite struct {
	lever
}

// NewEventBrite returns a new provider
func NewEventBrite() Provider {
	return &eventbrite{
		lever{
			"EventBrite",
			"eventbrite",
		},
	}
}
