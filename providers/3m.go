package providers

type _3M struct {
	myWorkdayJobs
}

func New3M() *_3M {
	return &_3M{
		myWorkdayJobs{
			"3M",
			"https://3m.wd1.myworkdayjobs.com",
			"/Search",
		},
	}
}
