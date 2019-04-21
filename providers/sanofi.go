package providers

type sanofi struct {
	students    myWorkdayJobs
	experienced myWorkdayJobs
}

func (sanofi *sanofi) RetrieveJobs(fn func(job *Job)) error {
	err := sanofi.experienced.RetrieveJobs(fn)
	if err != nil {
		return err
	}
	return sanofi.students.RetrieveJobs(fn)
}

// NewSanofi returns a new provider
func NewSanofi() Provider {
	return &sanofi{
		myWorkdayJobs{
			"Sanofi",
			"https://sanofi.wd3.myworkdayjobs.com",
			"/StudentPrograms",
		}, myWorkdayJobs{
			"Sanofi",
			"https://sanofi.wd3.myworkdayjobs.com",
			"/SanofiCareers",
		},
	}
}
