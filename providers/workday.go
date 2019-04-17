package providers

type workday struct {
	myWorkdayJobs
}

func NewWorkday() *workday {
	return &workday{
		myWorkdayJobs{
			"Workday",
			"https://workday.wd5.myworkdayjobs.com",
			"/Workday",
		},
	}
}
