# JobTracker

[![go report](https://goreportcard.com/badge/github.com/IDerr/jobtracker)](https://goreportcard.com/report/github.com/IDerr/jobtracker)
[![](https://godoc.org/github.com/IDerr/jobtracker/providers?status.svg)](https://godoc.org/github.com/IDerr/jobtracker/providers)

JobTracker aims to help you find your future dream job.
You can also use our library to scrap and export jobs from 150+ providers.

## Runnable software

### Compile

```shell
git clone https://github.com/IDerr/jobtracker.git
cd jobtracker
go get ./...
go build
```

### Notify

Jobtracker provides an easy way to track job offers.
Here is an example with Pushover. Discord and Pushbullet are also supported. See jobtracker help for more information

```shell
jobtracker notify pushover <token> <user> civiwebLatest.snap civiwebLatest
```

Just add this command to a crontab task, and voilà ! You'll get notified each time a job is added to the list

```shell
*/10 * * * * jobtracker notify pushover <token> <user> civiwebLatest.snap civiwebLatest
```

You can specify a list of providers to track multiple providers at the same time

```shell
jobtracker notify pushover <token> <user> civiwebLatest.snap civiwebLatest disney amazon netflix
```

### Print

The print command allows you to print the content of a snapshot or directly pulling data from any available provider.

```shell
jobtracker print [provider]...
```

### Filtering

You are also able to filter data from an external lua script.
Let's try with this 'test.lua' file

```lua
function filter(job)
	lower = string.lower(job.title)
	return string.match(lower, "data") == "data"
end
```

This command will then filter jobs from Disney and print only those whose title contains 'data'
```shell
jobtracker print -f=test.lua disney
```

## Library

### How to use
Example finding Netflix jobs

```go
p := providers.NewNetflix()	
p.RetrieveJobs(func(job *providers.Job) {	
  log.Println(job.Title, job.Company)	
})
```

### Providers and Companies

| Provider | Companies |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| 3M | 3M |
| Adobe | Adobe |
| Amazon | Amazon |
| Babylist | Babylist |
| Betclic | Betclic |
| Blizzard | Blizzard |
| Carta | Carta |
| Confluent | confluent |
| Coursera | Coursera |
| Dell | Dell |
| Disney | °O° Aulani, ABC Television, Disney*, ESPN, Marvel, Lucasfilm |
| Doctrine | Doctrine |
| DukeEnergy | Duke Energy |
| ERM | ERM |
| Eventbrite | Eventbrite |
| Fico | Fico |
| Flextronics | Flextronics |
| Gamestop | Gamestop |
| GumGum | Gum Gum |
| Hottopic | Hot Topic |
| Journy | Journy |
| Kering | Alexander McQueen, Balenciaga, Bottega Veneta, Boucheron, Brioni, Girard-Perregaux, Gucci, Kering, Pomellato, Qeelin, Saint Laurent, Ulysse Nardin |
| Kickstarter | Kickstarter |
| Lever | Lever |
| LinuxFoundation | Linux Foundation |
| Logitech | Astro, Blue Microphones, Jaybird, Logitech, Ultimate Ears |
| Mastercard | Mastercard |
| Medium | Medium |
| Netflix | Netflix |
| Nintendo | Nintendo |
| Npmjs | Npmjs |
| UniversityOfNevadaReno | University Of Nevada Reno |
| NYTimes | NYTimes |
| Oath | Oath, Yahoo, AOL, Engadget, Tumblr, Ryot, Makers, Verizon, BrightRoll, TechCrunch, HuffPost, Flurry |
| Outreach | Outreach |
| PaloAltoNetworks | Palo Alto Networks |
| Pokemon | The Pokémon Company |
| RollsRoyce | Rolls Royce |
| RosettaStone | Rosetta Stone |
| Salesforce | Salesforce |
| Samsung | Samsung |
| Sanofi | Sanofi |
| Scribd | Scribd |
| Soundcloud | Soundcloud |
| Strait | Strait |
| Thales | Thales |
| Trafigura | Trafigura, Puma Energy |
| Trainline | Trainline |
| Twitch | Twitch |
| Twitter | Twitter |
| UniversityOfChicago | The University Of Chicago |
| Vinted | Vinted |
| Whittard | Whittard |
| Workday | Workday |
| Github | Github |
| Lookout | Lookout |
| DigitalOcean | DigitalOcean |
| TripAdvisor | TripAdvisor |
| Fitbit | Fitbit |
| Airbnb | Airbnb |
| Evernote | Evernote |
| Twilio | Twilio |
| Pinterest | Pinterest |
| Vimeo | Vimeo |
| Surveymonkey | Surveymonkey |
| Docusign | Docusign |
| Casper | Casper |
| Metromile | Metromile |
| Squarespace | Squarespace |
| Nerdwallet | Nerdwallet |
| Buzzfeed | Buzzfeed |
| Thumbtack | Thumbtack |
| Expa | Expa |
| Carousell | Carousell |
| Classy | Classy |
| Thumbtack | Thumbtack |
| Lob | Lob |
| Handshake | Handshake |
| ClassPass | ClassPass |
| Genius | Genius |
| Curalate | Curalate |
| Qualtrics | Qualtrics |
| Envato | Envato |
| PocketGems | PocketGems |
| Ibotta | Ibotta |
| InterCom | InterCom |
| Massdrop | Massdrop |
| Gusto | Gusto |
| Payoff | Payoff |
| Granular | Granular |
| Zype | Zype |
| Smarkets | Smarkets |
| ImoDotIm | ImoDotIm |
| Wistia | Wistia |
| DoorDash | DoorDash |
| WarbyParker | WarbyParker |
| Strava | Strava |
| VirtaHealth | VirtaHealth |
| PureStorage | PureStorage |
| TrackMaven | TrackMaven |
| StackCommerce | StackCommerce |
| Patreon | Patreon |
| Twilio | Twilio |
| Signpost | Signpost |
| Bonobos | Bonobos |
| Box | Box |
| BritAndCo | BritAndCo |
| Kespry | Kespry |
| SimpleFinance | SimpleFinance |
| Postmates | Postmates |
| OmadaHealth | OmadaHealth |
| Chartboost | Chartboost |
| CourseHero | CourseHero |
| JWPlayer | JWPlayer |
| Mixpanel | Mixpanel |
| Carvana | Carvana |
| MongoDB | MongoDB |
| JauntVR | JauntVR |
| CommerceHub | CommerceHub |
| Unity | Unity |
| Spredfast | Spredfast |
| Quantifind | Quantifind |
| Justworks | Justworks |
| Splash | Splash |
| Magnetic | Magnetic |
| Agoda | Agoda |
| Eero | Eero |
| Netskope | Netskope |
| Blockchain | Blockchain |
| PDTPartners | PDTPartners |
| Rapid7 | Rapid7 |
| Giphy | Giphy |
| MagicLeap | MagicLeap |
| AppLovin | AppLovin |
| ShipHawk | ShipHawk |
| TripAdvisor | TripAdvisor |
| Pindrop | Pindrop |
| Takealot | Takealot |
| TheWorkingGroup | TheWorkingGroup |
| TheSourcery | TheSourcery |
| CampaignMonitor | CampaignMonitor |
| RallyHealth | RallyHealth |
| Persado | Persado |
| VaynerMedia | VaynerMedia |
| Pager | Pager |
| Picarro | Picarro |
| Beekeeper | Beekeeper |
| Booking | Booking |
| Noom | Noom |
| DigitalOcean | DigitalOcean |
| Skookum | Skookum |
| Helix | Helix |
| Malwarebytes | Malwarebytes |
| Braintree | Braintree |
| Pager | Pager |
| Piktochart | Piktochart |
| Coupang | Coupang |
| Zanbato | Zanbato |
| Current | Current |
| Embark | Embark |
| Climb | Climb |
| TrueMotion | TrueMotion |
| Teralytics | Teralytics |
| Oseberg | Oseberg |
| Peek | Peek |
| Civiweb | All "VI" jobs, French citizens only |
