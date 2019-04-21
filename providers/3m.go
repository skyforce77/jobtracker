package providers

type _3M struct {
	myWorkdayJobs
}

// New3M returns a new provider
func New3M() Provider {
	return &_3M{
		myWorkdayJobs{
			"3M",
			"https://3m.wd1.myworkdayjobs.com",
			"/Search",
		},
	}
}
