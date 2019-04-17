package providers

type flextronics struct {
	myWorkdayJobs
}

func NewFlextronics() *flextronics {
	return &flextronics{
		myWorkdayJobs{
			"Flextronics",
			"https://flextronics.wd1.myworkdayjobs.com",
			"/Careers",
		},
	}
}
