package providers

type mastercard struct {
	myWorkdayJobs
}

// NewMastercard returns a new provider
func NewMastercard() Provider {
	return &mastercard{
		myWorkdayJobs{
			"Mastercard",
			"https://mastercard.wd1.myworkdayjobs.com",
			"/CorporateCareers",
		},
	}
}
