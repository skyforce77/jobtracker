package providers

type surveymonkey struct {
	greenhouse
}

// NewSurveymonkey returns a new provider
func NewSurveymonkey() Provider {
	return &surveymonkey{
		greenhouse{
			"Surveymonkey",
			"surveymonkey",
		},
	}
}
