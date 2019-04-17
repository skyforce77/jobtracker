package providers

type nshe struct {
	myWorkdayJobs
}

func NewUniversityOfNevadaReno() *nshe {
	return &nshe{
		myWorkdayJobs{
			"University Of Nevada, Reno",
			"https://nshe.wd1.myworkdayjobs.com",
			"/UNR-external",
		},
	}
}
