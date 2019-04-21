package providers

type talkiq struct {
	greenhouse
}

// NewTalkIQ returns a new provider
func NewTalkIQ() Provider {
	return &talkiq{
		greenhouse{
			"TalkIQ",
			"talkiq",
		},
	}
}
