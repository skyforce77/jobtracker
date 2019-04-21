package providers

type dell struct {
	myWorkdayJobs
}

// NewDell returns a new provider
func NewDell() Provider {
	return &dell{
		myWorkdayJobs{
			"Dell",
			"https://dell.wd1.myworkdayjobs.com",
			"/External",
		},
	}
}
