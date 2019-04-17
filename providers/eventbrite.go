package providers

type eventbrite struct {
	lever
}

func NewEventBrite() *eventbrite {
	return &eventbrite{
		lever{
			"EventBrite",
			"eventbrite",
		},
	}
}
