package providers

type oath struct {
	myWorkdayJobs
}

// NewOath returns a new provider
func NewOath() Provider {
	return &oath{
		myWorkdayJobs{
			"Oath",
			"https://oath.wd5.myworkdayjobs.com",
			"/careers",
		},
	}
}
