package providers

type gamestop struct {
	myWorkdayJobs
}

// NewGamestop returns a new provider
func NewGamestop() Provider {
	return &gamestop{
		myWorkdayJobs{
			"Gamestop",
			"https://gamestop.wd5.myworkdayjobs.com",
			"/Store",
		},
	}
}
