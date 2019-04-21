package providers

type fico struct {
	myWorkdayJobs
}

// NewFico returns a new provider
func NewFico() Provider {
	return &fico{
		myWorkdayJobs{
			"Fico",
			"https://fico.wd1.myworkdayjobs.com",
			"/External",
		},
	}
}
