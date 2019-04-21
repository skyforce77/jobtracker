package providers

type rollsRoyce struct {
	myWorkdayJobs
}

// NewRollsRoyce returns a new provider
func NewRollsRoyce() Provider {
	return &rollsRoyce{
		myWorkdayJobs{
			"Rolls-Royce",
			"https://rollsroyce.wd3.myworkdayjobs.com",
			"/professional",
		},
	}
}
