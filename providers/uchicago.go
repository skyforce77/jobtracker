package providers

type uchicago struct {
	myWorkdayJobs
}

func NewUniversityOfChicago() *uchicago {
	return &uchicago{
		myWorkdayJobs{
			"The University Of Chicago",
			"https://uchicago.wd5.myworkdayjobs.com",
			"/External",
		},
	}
}
