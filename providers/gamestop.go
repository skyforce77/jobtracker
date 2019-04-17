package providers

type gamestop struct {
	myWorkdayJobs
}

func NewGamestop() *gamestop {
	return &gamestop{
		myWorkdayJobs{
			"Gamestop",
			"https://gamestop.wd5.myworkdayjobs.com",
			"/Store",
		},
	}
}
