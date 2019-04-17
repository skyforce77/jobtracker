package providers

type dell struct {
	myWorkdayJobs
}

func NewDell() *dell {
	return &dell{
		myWorkdayJobs{
			"Dell",
			"https://dell.wd1.myworkdayjobs.com",
			"/External",
		},
	}
}
