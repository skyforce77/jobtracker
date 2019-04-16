package providers

type sanofi struct {
	myWorkdayJobs
}

func NewSanofi() *sanofi {
	return &sanofi{
		myWorkdayJobs{
			"Sanofi",
			"https://sanofi.wd3.myworkdayjobs.com",
			"/StudentPrograms",
		},
	}
}
