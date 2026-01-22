package providers

import (
	"bytes"
	"container/list"
	"crypto/md5"
	"reflect"
)

// JobType represents a job type, by schedule and by contract
type JobType string

// Job type constants
const (
	FullTime       JobType = "Full time"
	PartTime       JobType = "Part time"
	FixedTerm      JobType = "Fixed term"
	Temporary      JobType = "Temporary"
	Internship     JobType = "Internship"
	Apprenticeship JobType = "Apprenticeship"
	VI             JobType = "Volontariat International (French Citizen)"
)

var (
	providers = []Provider{
		NewOpenAI(),
		NewNotion(),
		NewRamp(),
		NewLinear(),
		NewDeel(),
		NewGreenhouse(),
		NewAirbnb(),
		NewTwitch(),
		NewTwilio(),
		NewPinterest(),
		NewMongoDB(),
		NewInterCom(),
		NewSquarespace(),
		NewFigma(),
		NewStripe(),
		NewDiscord(),
		NewCoinbase(),
		NewAnthropic(),
		NewNetflix(),
		NewAdobe(),
		NewAmazon(),
		NewBlizzard(),
		NewConfluent(),
		NewCoursera(),
		NewDisney(),
		NewKickStarter(),
		NewLinuxFoundation(),
		NewLogitech(),
		NewNintendo(),
		NewSalesforce(),
		NewUnity(),
		NewDigitalOcean(),
		NewDocusign(),
		NewNerdwallet(),
		NewBuzzfeed(),
		NewQualtrics(),
		NewIbotta(),
		NewStrava(),
		NewPureStorage(),
		NewBox(),
		NewCourseHero(),
		NewMixpanel(),
		NewRapid7(),
		NewMalwarebytes(),
		NewLever(),
		NewSpotify(),
	}
)

// Job is a standardized job offer
type Job struct {
	// Title specifies the job title
	Title string `json:"title"`

	// Company specifies the company offering the job
	//
	// A single provider may provide multiple companies
	Company string `json:"company"`

	// Location is the location of the job's office
	Location string `json:"location"`

	// Type specifies the job schedule or contract type
	Type string `json:"type"`

	// Desc is the job's description
	Desc string `json:"description"`

	// Link refers to an HTTP URL providing the job offer
	Link string `json:"link"`

	// Misc may contain more specific information
	Misc map[string]string `json:"misc"`
}

// Hash returns the structure md5 sum based on job title, company and location
func (job *Job) Hash() [16]byte {
	var b bytes.Buffer
	b.Write([]byte(job.Title))
	b.Write([]byte(job.Company))
	b.Write([]byte(job.Location))
	b.Write([]byte(job.Link))
	return md5.Sum(b.Bytes())
}

// Provider is able to scrap jobs from a specific website
type Provider interface {
	// RetrieveJobs starts the jobs scraping
	//
	// This call is blocking and calling the function argument synchronously
	RetrieveJobs(func(job *Job)) error
}

// ApplicationRequest contains the data needed to apply for a job
type ApplicationRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone,omitempty"`
	Resume       []byte `json:"resume,omitempty"`
	ResumeURL    string `json:"resume_url,omitempty"`
	CoverLetter  string `json:"cover_letter,omitempty"`
	LinkedIn     string `json:"linkedin,omitempty"`
	Website      string `json:"website,omitempty"`
	Location     string `json:"location,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// ApplicationResult contains the response from a job application
type ApplicationResult struct {
	Success       bool   `json:"success"`
	ApplicationID string `json:"application_id,omitempty"`
	Message       string `json:"message,omitempty"`
	Error         string `json:"error,omitempty"`
}

// Applicant is an optional interface that providers can implement
// to support automatic job applications
type Applicant interface {
	ApplyToJob(jobURL string, req *ApplicationRequest) (*ApplicationResult, error)
}

// Collect helps you recovering a list of jobs from a Provider
func Collect(provider Provider) *list.List {
	lst := list.New()
	provider.RetrieveJobs(func(job *Job) {
		lst.PushBack(job)
	})
	return lst
}

// IterateOver make iterating over a job list easy
//
// You may want to use it after Collect
func IterateOver(lst *list.List, fn func(*Job)) {
	n := lst.Front()
	for n != nil {
		v := n.Value.(*Job)
		fn(v)
		n = n.Next()
	}
}

// RetrieveAsync calls your function argument asynchronously
//
// This call is blocking to let you know when it finishes
//
// You should use it in a goroutine to avoid blocking
func RetrieveAsync(provider Provider, fn func(*Job)) {
	provider.RetrieveJobs(func(job *Job) {
		go fn(job)
	})
}

// GetProviders return a complete list of available providers
func GetProviders() []Provider {
	return providers
}

// Diff represents a diff between two providers
type Diff struct {
	Added   []*Job
	Removed []*Job
}

// NewDiff creates a diff
func NewDiff(provider Provider, provider2 Provider) (*Diff, error) {
	m := make(map[[16]byte]*Job)
	added := list.New()

	err := provider.RetrieveJobs(func(job *Job) {
		m[job.Hash()] = job
	})
	if err != nil {
		return nil, err
	}

	err = provider2.RetrieveJobs(func(job *Job) {
		hash := job.Hash()
		if _, ok := m[hash]; ok {
			delete(m, hash)
		} else {
			added.PushBack(job)
		}
	})
	if err != nil {
		return nil, err
	}

	diff := &Diff{}

	diff.Removed = make([]*Job, len(m))
	i := 0
	for _, v := range m {
		diff.Removed[i] = v
		i++
	}

	diff.Added = make([]*Job, added.Len())
	i = 0
	IterateOver(added, func(job *Job) {
		diff.Added[i] = job
		i++
	})

	return diff, nil
}

// ProviderFromName returns a provider from its name
func ProviderFromName(name string) Provider {
	for _, p := range GetProviders() {
		typ := reflect.TypeOf(p)
		nam := typ.String()[11:]
		if nam[0] == '_' {
			nam = nam[1:]
		}

		if nam == name {
			return p
		}
	}
	return nil
}

