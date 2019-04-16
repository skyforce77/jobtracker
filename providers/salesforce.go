package providers

type salesforce struct {
	myWorkdayJobs
}

func NewSalesforce() *salesforce {
	return &salesforce{
		myWorkdayJobs{
			"Salesforce",
			"https://salesforce.wd1.myworkdayjobs.com",
			"/External_Career_Site",
		},
	}
}
