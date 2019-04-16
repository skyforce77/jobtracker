package providers

type fico struct {
	myWorkdayJobs
}

func NewFico() *fico {
	return &fico{
		myWorkdayJobs{
			"Fico",
			"https://fico.wd1.myworkdayjobs.com",
			"/External",
		},
	}
}
