# JobTracker

![go report](https://goreportcard.com/badge/github.com/IDerr/jobtracker)

JobTracker aims to help you find your future dream job.
You can use our library to scrap and export jobs from 50+ providers.

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
