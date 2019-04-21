package providers

type nytimes struct {
	myWorkdayJobs
}

// NewNYTimes returns a new provider
func NewNYTimes() Provider {
	return &nytimes{
		myWorkdayJobs{
			"New York Times",
			"https://nytimes.wd5.myworkdayjobs.com",
			"/News",
		},
	}
}
