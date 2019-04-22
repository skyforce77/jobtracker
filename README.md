# JobTracker

[![go report](https://goreportcard.com/badge/github.com/IDerr/jobtracker)](https://goreportcard.com/report/github.com/IDerr/jobtracker)
[![](https://godoc.org/github.com/IDerr/jobtracker/providers?status.svg)](https://godoc.org/github.com/IDerr/jobtracker/providers)

JobTracker aims to help you find your future dream job.
You can use our library to scrap and export jobs from 150+ providers.

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
