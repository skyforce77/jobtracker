package providers

type erm struct {
	myWorkdayJobs
}

func NewERM() *erm {
	return &erm{
		myWorkdayJobs{
			"ERM",
			"https://erm.wd3.myworkdayjobs.com",
			"/ERM_Careers",
		},
	}
}
