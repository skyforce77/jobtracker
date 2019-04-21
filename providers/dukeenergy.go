package providers

type dukeEnergy struct {
	myWorkdayJobs
}

// NewDukeEnergy returns a new provider
func NewDukeEnergy() Provider {
	return &dukeEnergy{
		myWorkdayJobs{
			"Duke Energy",
			"https://dukeenergy.wd1.myworkdayjobs.com",
			"/search",
		},
	}
}
