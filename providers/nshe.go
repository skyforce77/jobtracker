package providers

type nshe struct {
	myWorkdayJobs
}

// NewUniversityOfNevadaReno returns a new provider
func NewUniversityOfNevadaReno() Provider {
	return &nshe{
		myWorkdayJobs{
			"University Of Nevada, Reno",
			"https://nshe.wd1.myworkdayjobs.com",
			"/UNR-external",
		},
	}
}
