package providers

type thales struct {
	myWorkdayJobs
}

// NewThales returns a new provider
func NewThales() Provider {
	return &thales{
		myWorkdayJobs{
			"Thales",
			"https://thales.wd3.myworkdayjobs.com",
			"/Careers",
		},
	}
}
