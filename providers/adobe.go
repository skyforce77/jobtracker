package providers

type adobe struct {
	experienced myWorkdayJobs
	university  myWorkdayJobs
}

func (adobe *adobe) RetrieveJobs(fn func(job *Job)) error {
	err := adobe.experienced.RetrieveJobs(fn)
	if err != nil {
		return err
	}
	return adobe.university.RetrieveJobs(fn)
}

// NewAdobe returns a new provider
func NewAdobe() Provider {
	return &adobe{
		myWorkdayJobs{
			"Adobe",
			"https://adobe.wd5.myworkdayjobs.com",
			"/external_experienced",
		},
		myWorkdayJobs{
			"Adobe",
			"https://adobe.wd5.myworkdayjobs.com",
			"/external_university",
		},
	}
}
