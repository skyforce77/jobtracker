package providers

type erm struct {
	myWorkdayJobs
}

// NewERM returns a new provider
func NewERM() Provider {
	return &erm{
		myWorkdayJobs{
			"ERM",
			"https://erm.wd3.myworkdayjobs.com",
			"/ERM_Careers",
		},
	}
}
