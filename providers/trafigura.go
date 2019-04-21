package providers

type trafigura struct {
	myWorkdayJobs
}

// NewTrafigura returns a new provider
func NewTrafigura() Provider {
	return &trafigura{
		myWorkdayJobs{
			"Trafigura",
			"https://trafigura.wd3.myworkdayjobs.com",
			"/Puma_Energy_Careers",
		},
	}
}
