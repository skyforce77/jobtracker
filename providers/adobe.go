package providers

type adobe struct {
	myWorkdayJobs
}

func NewAdobe() *adobe {
	return &adobe{
		myWorkdayJobs{
			"Adobe",
			"https://adobe.wd5.myworkdayjobs.com",
			"/external_experienced",
		},
	}
}