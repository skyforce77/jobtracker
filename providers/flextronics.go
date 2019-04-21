package providers

type flextronics struct {
	myWorkdayJobs
}

// NewFlextronics returns a new provider
func NewFlextronics() Provider {
	return &flextronics{
		myWorkdayJobs{
			"Flextronics",
			"https://flextronics.wd1.myworkdayjobs.com",
			"/Careers",
		},
	}
}
