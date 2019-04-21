package providers

type jwplayer struct {
	greenhouse
}

// NewJWPlayer returns a new provider
func NewJWPlayer() Provider {
	return &jwplayer{
		greenhouse{
			"JWPlayer",
			"jwplayer",
		},
	}
}
