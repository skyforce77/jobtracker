package providers

type samsung struct {
	jobVite       jobVite
	myWorkdayJobs myWorkdayJobs
}

func (samsung *samsung) RetrieveJobs(fn func(job *Job)) {
	samsung.myWorkdayJobs.RetrieveJobs(fn)
	samsung.jobVite.RetrieveJobs(fn)
}

func NewSamsung() *samsung {
	return &samsung{
		jobVite{
			"Samsung",
			"http://jobs.jobvite.com/samsungssi/jobs",
		},
		myWorkdayJobs{
			"Samsung",
			"https://sec.wd3.myworkdayjobs.com",
			"/Samsung_Careers",
		},
	}
}
