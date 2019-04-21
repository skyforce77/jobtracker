package providers

type workday struct {
	myWorkdayJobs
}

// NewWorkday returns a new provider
func NewWorkday() Provider {
	return &workday{
		myWorkdayJobs{
			"Workday",
			"https://workday.wd5.myworkdayjobs.com",
			"/Workday",
		},
	}
}
