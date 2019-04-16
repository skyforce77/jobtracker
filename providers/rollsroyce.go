package providers

type rollsRoyce struct {
	myWorkdayJobs
}

func NewRollsRoyce() *rollsRoyce {
	return &rollsRoyce{
		myWorkdayJobs{
			"Rolls-Royce",
			"https://rollsroyce.wd3.myworkdayjobs.com",
			"/professional",
		},
	}
}
