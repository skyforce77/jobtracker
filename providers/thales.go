package providers

type thales struct {
	myWorkdayJobs
}

func NewThales() *thales {
	return &thales{
		myWorkdayJobs{
			"Thales",
			"https://thales.wd3.myworkdayjobs.com",
			"/Careers",
		},
	}
}
