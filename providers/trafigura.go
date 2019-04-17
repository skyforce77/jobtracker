package providers

type trafigura struct {
	myWorkdayJobs
}

func NewTrafigura() *trafigura {
	return &trafigura{
		myWorkdayJobs{
			"Trafigura",
			"https://trafigura.wd3.myworkdayjobs.com",
			"/Puma_Energy_Careers",
		},
	}
}
