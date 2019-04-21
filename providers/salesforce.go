package providers

type salesforce struct {
	myWorkdayJobs
}

// NewSalesforce returns a new provider
func NewSalesforce() Provider {
	return &salesforce{
		myWorkdayJobs{
			"Salesforce",
			"https://salesforce.wd1.myworkdayjobs.com",
			"/External_Career_Site",
		},
	}
}
