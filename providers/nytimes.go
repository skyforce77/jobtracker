package providers

type nytimes struct {
	myWorkdayJobs
}

func NewNYTimes() *nytimes {
	return &nytimes{
		myWorkdayJobs{
			"New York Times",
			"https://nytimes.wd5.myworkdayjobs.com",
			"/News",
		},
	}
}
