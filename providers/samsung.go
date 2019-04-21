package providers

type samsung struct {
	jobVite       jobVite
	myWorkdayJobs myWorkdayJobs
}

func (samsung *samsung) RetrieveJobs(fn func(job *Job)) error {
	err := samsung.myWorkdayJobs.RetrieveJobs(fn)
	if err != nil {
		return err
	}
	return samsung.jobVite.RetrieveJobs(fn)
}

// NewSamsung returns a new provider
func NewSamsung() Provider {
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
