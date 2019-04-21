package providers

type docusign struct {
	greenhouse
}

// NewDocusign returns a new provider
func NewDocusign() Provider {
	return &docusign{
		greenhouse{
			"Docusign",
			"docusign",
		},
	}
}
