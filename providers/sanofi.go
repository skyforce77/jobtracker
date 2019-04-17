package providers

type sanofi struct {
	students    myWorkdayJobs
	experienced myWorkdayJobs
}

func (sanofi *sanofi) RetrieveJobs(fn func(job *Job)) {
	sanofi.experienced.RetrieveJobs(fn)
	sanofi.students.RetrieveJobs(fn)
}

func NewSanofi() *sanofi {
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
