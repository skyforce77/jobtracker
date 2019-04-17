package providers

type dukeEnergy struct {
	myWorkdayJobs
}

func NewDukeEnergy() *dukeEnergy {
	return &dukeEnergy{
		myWorkdayJobs{
			"Duke Energy",
			"https://dukeenergy.wd1.myworkdayjobs.com",
			"/search",
		},
	}
}
