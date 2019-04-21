package providers

type workingatbooking struct {
	greenhouse
}

// NewBooking returns a new provider
func NewBooking() Provider {
	return &workingatbooking{
		greenhouse{
			"Booking",
			"workingatbooking",
		},
	}
}
