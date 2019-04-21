package providers

type postmates struct {
	greenhouse
}

// NewPostmates returns a new provider
func NewPostmates() Provider {
	return &postmates{
		greenhouse{
			"Postmates",
			"postmates",
		},
	}
}
