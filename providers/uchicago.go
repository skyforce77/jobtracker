package providers

type uchicago struct {
	myWorkdayJobs
}

// NewUniversityOfChicago returns a new provider
func NewUniversityOfChicago() Provider {
	return &uchicago{
		myWorkdayJobs{
			"The University Of Chicago",
			"https://uchicago.wd5.myworkdayjobs.com",
			"/External",
		},
	}
}
