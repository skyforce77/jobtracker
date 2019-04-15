package providers

type JobType string

const (
	FullTime       JobType = "Full time"
	PartTime       JobType = "Part time"
	Internship     JobType = "Internship"
	Apprenticeship JobType = "Apprenticeship"
)

type Job struct {
	Title    string
	Company  string
	Location string
	Type     string
	Desc     string
	Link     string
	Misc     map[string]string
}

type Provider interface {
	ListJobs() []*Job
}
