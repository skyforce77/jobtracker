package providers

type evernote struct {
	greenhouse
}

// NewEvernote returns a new provider
func NewEvernote() Provider {
	return &evernote{
		greenhouse{
			"Evernote",
			"evernote",
		},
	}
}
