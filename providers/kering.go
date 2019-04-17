package providers

type kering struct {
	myWorkdayJobs
}

func NewKering() *kering {
	return &kering{
		myWorkdayJobs{
			"Kering",
			"https://kering.wd3.myworkdayjobs.com",
			"/Kering",
		},
	}
}
