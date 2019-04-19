package providers

import (
	"container/list"
)

type JobType string

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

type Job struct {
	Title    string            `json:"title"`
	Company  string            `json:"company"`
	Location string            `json:"location"`
	Type     string            `json:"type"`
	Desc     string            `json:"description"`
	Link     string            `json:"link"`
	Misc     map[string]string `json:"misc"`
}

type Provider interface {
	RetrieveJobs(func(job *Job))
}

func Collect(provider Provider) *list.List {
	lst := list.New()
	provider.RetrieveJobs(func(job *Job) {
		lst.PushBack(job)
	})
	return lst
}

func IterateOver(lst *list.List, fn func(*Job)) {
	n := lst.Front()
	for n != nil {
		v := n.Value.(*Job)
		fn(v)
		n = n.Next()
	}
}

func RetrieveAsync(provider Provider, fn func(*Job)) {
	provider.RetrieveJobs(func(job *Job) {
		go fn(job)
	})
}

func GetProviders() []Provider {
	return providers
}
