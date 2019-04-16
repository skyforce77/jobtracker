package providers

type mastercard struct {
	myWorkdayJobs
}

func NewMastercard() *mastercard {
	return &mastercard{
		myWorkdayJobs{
			"Mastercard",
			"https://mastercard.wd1.myworkdayjobs.com",
			"/CorporateCareers",
		},
	}
}
