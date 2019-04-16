package providers

type oath struct {
	myWorkdayJobs
}

func NewOath() *oath {
	return &oath{
		myWorkdayJobs{
			"Oath",
			"https://oath.wd5.myworkdayjobs.com",
			"/careers",
		},
	}
}
