package providers

type kering struct {
	myWorkdayJobs
}

// NewKering returns a new provider
func NewKering() Provider {
	return &kering{
		myWorkdayJobs{
			"Kering",
			"https://kering.wd3.myworkdayjobs.com",
			"/Kering",
		},
	}
}
