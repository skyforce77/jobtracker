package providers

import (
	"container/list"
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
)

var (
	providers = []Provider{
		New3M(),
		NewAdobe(),
		NewAmazon(),
		NewBabylist(),
		NewBetclic(),
		NewBlizzard(),
		NewCarta(),
		NewConfluent(),
		NewCoursera(),
		NewDell(),
		NewDisney(),
		NewDoctrine(),
		NewDukeEnergy(),
		NewERM(),
		NewEventBrite(),
		NewFico(),
		NewFlextronics(),
		NewGamestop(),
		NewGumGum(),
		NewHottopic(),
		NewJourny(),
		NewKering(),
		NewKickStarter(),
		NewLever(),
		NewLinuxFoundation(),
		NewLogitech(),
		NewMastercard(),
		NewMedium(),
		NewNetflix(),
		NewNintendo(),
		NewNpmjs(),
		NewUniversityOfNevadaReno(),
		NewNYTimes(),
		NewOath(),
		NewOutreach(),
		NewPaloAltoNetworks(),
		NewPokemon(),
		NewRollsRoyce(),
		NewRosettaStone(),
		NewSalesforce(),
		NewSamsung(),
		NewSanofi(),
		NewScribd(),
		NewSoundcloud(),
		NewStrait(),
		NewThales(),
		NewTrafigura(),
		NewTrainline(),
		NewTwitch(),
		NewTwitter(),
		NewUniversityOfChicago(),
		NewVinted(),
		NewWhittard(),
		NewWorkday(),
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

// Provider is able to scrap jobs from a specific website
type Provider interface {
	// RetrieveJobs starts the jobs scraping
	//
	// This call is blocking and calling the function argument synchronously
	RetrieveJobs(func(job *Job)) error
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
