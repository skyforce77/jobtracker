package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type myWorkdayJobs struct {
	company string
	url     string
	home    string
}

type myWorkdayJobsPage struct {
	ID     string `json:"id"`
	Widget string `json:"widget"`
	Body   struct {
		ID       string `json:"id"`
		Label    string `json:"label"`
		Widget   string `json:"widget"`
		Children []struct {
			ID             string `json:"id"`
			Widget         string `json:"widget"`
			Ecid           string `json:"ecid"`
			FacetContainer struct {
				ID         string `json:"id"`
				Widget     string `json:"widget"`
				Ecid       string `json:"ecid"`
				SearchText struct {
					Widget string `json:"widget"`
				} `json:"searchText"`
				PaginationCount struct {
					ID          string `json:"id"`
					Widget      string `json:"widget"`
					Ecid        string `json:"ecid"`
					Value       int    `json:"value"`
					TotalDigits int    `json:"totalDigits"`
					Precision   int    `json:"precision"`
				} `json:"paginationCount"`
				Offset struct {
					ID          string `json:"id"`
					Widget      string `json:"widget"`
					Ecid        string `json:"ecid"`
					Value       int    `json:"value"`
					TotalDigits int    `json:"totalDigits"`
					Precision   int    `json:"precision"`
				} `json:"offset"`
				Facets []struct {
					ID          string `json:"id"`
					Label       string `json:"label"`
					Widget      string `json:"widget"`
					Ecid        string `json:"ecid"`
					Iid         string `json:"iid"`
					FacetValues []struct {
						ID     string `json:"id"`
						Label  string `json:"label"`
						Widget string `json:"widget"`
						Ecid   string `json:"ecid"`
						Iid    string `json:"iid"`
						Count  int    `json:"count"`
					} `json:"facetValues"`
					ExpandState string `json:"expandState"`
				} `json:"facets"`
			} `json:"facetContainer,omitempty"`
			Children []struct {
				ID        string `json:"id"`
				Widget    string `json:"widget"`
				Ecid      string `json:"ecid"`
				ListItems []struct {
					ID     string `json:"id"`
					Widget string `json:"widget"`
					Ecid   string `json:"ecid"`
					Title  struct {
						ID           string `json:"id"`
						Widget       string `json:"widget"`
						Ecid         string `json:"ecid"`
						PropertyName string `json:"propertyName"`
						Singular     bool   `json:"singular"`
						Instances    []struct {
							ID     string `json:"id"`
							Widget string `json:"widget"`
							Text   string `json:"text"`
							Action string `json:"action"`
							V      bool   `json:"v"`
						} `json:"instances"`
						CommandLink string `json:"commandLink"`
						MultiSelect bool   `json:"multiSelect"`
					} `json:"title"`
					Subtitles []struct {
						ID        string `json:"id"`
						Widget    string `json:"widget"`
						Ecid      string `json:"ecid"`
						Instances []struct {
							ID     string `json:"id"`
							Widget string `json:"widget"`
							Text   string `json:"text"`
						} `json:"instances"`
						MultiSelect bool `json:"multiSelect"`
					} `json:"subtitles"`
				} `json:"listItems"`
				Children []struct {
					ID         string `json:"id"`
					Widget     string `json:"widget"`
					Ecid       string `json:"ecid"`
					IconName   string `json:"iconName,omitempty"`
					ImageLabel string `json:"imageLabel,omitempty"`
					Children   []struct {
						ID     string `json:"id"`
						Label  string `json:"label"`
						Widget string `json:"widget"`
						Ecid   string `json:"ecid"`
						Values []struct {
							Label      string `json:"label"`
							URI        string `json:"uri"`
							HTTPMethod string `json:"httpMethod"`
							Enabled    bool   `json:"enabled"`
						} `json:"values,omitempty"`
						URI               string `json:"uri"`
						FooterOrder       int    `json:"footerOrder,omitempty"`
						HTTPMethod        string `json:"httpMethod,omitempty"`
						Rule              int    `json:"rule,omitempty"`
						JobCode           string `json:"jobCode,omitempty"`
						ContextID         string `json:"contextId,omitempty"`
						APIKey            string `json:"apiKey,omitempty"`
						DataVariableName  string `json:"dataVariableName,omitempty"`
						GadgetRedirectURL string `json:"gadgetRedirectUrl,omitempty"`
						DataMode          string `json:"dataMode,omitempty"`
						ClientDataURL     string `json:"clientDataUrl,omitempty"`
					} `json:"children,omitempty"`
					PropertyName string `json:"propertyName,omitempty"`
					Text         string `json:"text,omitempty"`
					Value        string `json:"value,omitempty"`
				} `json:"children"`
			} `json:"children,omitempty"`
			EndPoints []struct {
				Type string `json:"type"`
				URI  string `json:"uri"`
				ID   string `json:"id"`
			} `json:"endPoints,omitempty"`
			UILabelIdentifier string `json:"uiLabelIdentifier,omitempty"`
		} `json:"children"`
		TabContent bool `json:"tabContent"`
	} `json:"body"`
	CurrentTime       int    `json:"currentTime"`
	Ecid              string `json:"ecid"`
	FailedSignonCount int    `json:"failedSignonCount"`
	Header            struct {
		ID             string `json:"id"`
		Label          string `json:"label"`
		Widget         string `json:"widget"`
		LanguagePicker struct {
			ID     string `json:"id"`
			Label  string `json:"label"`
			Widget string `json:"widget"`
			Ecid   string `json:"ecid"`
			Values []struct {
				Label      string `json:"label"`
				URI        string `json:"uri"`
				HTTPMethod string `json:"httpMethod"`
				Enabled    bool   `json:"enabled"`
				Type       string `json:"type"`
			} `json:"values"`
			FooterOrder int `json:"footerOrder"`
			Rule        int `json:"rule"`
		} `json:"languagePicker"`
		Authentication struct {
			ID                                     string `json:"id"`
			Widget                                 string `json:"widget"`
			SignInRequestURI                       string `json:"signInRequestUri"`
			ForgotPasswordRequestURI               string `json:"forgotPasswordRequestUri"`
			CreateAccountRequestURI                string `json:"createAccountRequestUri"`
			ResetPasswordRequestURI                string `json:"resetPasswordRequestUri"`
			ResendAccountActivationEmailRequestURI string `json:"resendAccountActivationEmailRequestUri"`
			LogoURL                                string `json:"logoUrl"`
			Title                                  string `json:"title"`
			ShowAsLink                             bool   `json:"showAsLink"`
			State                                  string `json:"state"`
			SignOutURI                             string `json:"signOutUri"`
			SuccessRedirectURI                     string `json:"successRedirectUri"`
			UILabelIdentifier                      string `json:"uiLabelIdentifier"`
		} `json:"authentication"`
		HeaderFashion      string `json:"headerFashion"`
		TitleAlignment     string `json:"titleAlignment"`
		BackgroundImageURL string `json:"backgroundImageUrl"`
	} `json:"header"`
	Title struct {
		ID     string `json:"id"`
		Widget string `json:"widget"`
		Text   string `json:"text"`
	} `json:"title"`
	WindowTitle      string `json:"windowTitle"`
	HomeURI          string `json:"homeUri"`
	LocaleAttributes struct {
		DateOrder            string `json:"dateOrder"`
		DateToolTip          string `json:"dateToolTip"`
		DecimalSeparator     string `json:"decimalSeparator"`
		FirstDayOfWeek       string `json:"firstDayOfWeek"`
		HourClock            string `json:"hourClock"`
		ThousandsSeparator   string `json:"thousandsSeparator"`
		UserLocale           string `json:"userLocale"`
		ID                   string `json:"id"`
		WorkdayDayMonthRange string `json:"workdayDayMonthRange"`
		WorkdayFullDateRange string `json:"workdayFullDateRange"`
	} `json:"localeAttributes"`
	OpenGraphAttributes struct {
		Type        string `json:"type"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		ImageURL    string `json:"imageUrl"`
		Description string `json:"description"`
		ID          string `json:"id"`
	} `json:"openGraphAttributes"`
	Mode          string `json:"mode"`
	Notifications struct {
		Widget            string `json:"widget"`
		RetrieveUnreadURI string `json:"retrieveUnreadUri"`
		RetrieveUnseenURI string `json:"retrieveUnseenUri"`
		TotalCount        int    `json:"totalCount"`
		UnreadCount       int    `json:"unreadCount"`
		UnseenCount       int    `json:"unseenCount"`
	} `json:"notifications"`
	RequestURI             string `json:"requestUri"`
	SessionSecureToken     string `json:"sessionSecureToken"`
	UserLanguageCode       string `json:"userLanguageCode"`
	PreferredLabelPosition string `json:"preferredLabelPosition"`
	NotificationAlertURI   string `json:"notificationAlertUri"`
	PmdFetchTime           int    `json:"pmdFetchTime"`
	SmdFetchTime           int    `json:"smdFetchTime"`
}

func (mwj *myWorkdayJobs) readPage(index int, mwjp *myWorkdayJobsPage, fn func(job *Job)) error {
	client := &http.Client{}

	lnItems := 0
	for _, c := range mwjp.Body.Children {
		if c.Widget == "facetSearchResult" {
			for _, w := range c.Children {
				if w.Widget == "facetSearchResultList" {
					lnItems = len(w.ListItems)
					for _, item := range w.ListItems {
						job := Job{
							Title:    item.Title.Instances[0].Text,
							Company:  mwj.company,
							Location: item.Subtitles[0].Instances[0].Text,
							Link:     mwj.url + item.Title.CommandLink,
							Type:     string(FullTime),
						}

						req, err := http.NewRequest("GET", mwj.url+item.Title.CommandLink, nil)
						if err != nil {
							return err
						}
						req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
						req.Header.Add("Accept", "application/json,application/xml")
						res, err := client.Do(req)
						if err != nil {
							return err
						}

						if res.StatusCode != 200 {
							return HandleStatus(res)
						}

						body, err := ioutil.ReadAll(res.Body)
						if err != nil {
							return err
						}
						res.Body.Close()

						page := myWorkdayJobsPage{}
						err = json.Unmarshal(body, &page)
						if err != nil {
							return err
						}

						job.Desc = page.Body.Children[1].Children[0].Children[2].Text
						fn(&job)
					}
				}
			}

			for _, point := range c.EndPoints {
				if point.Type == "Pagination" && lnItems != 0 {
					url := mwj.url + point.URI + "/" + strconv.Itoa(index+lnItems)
					req, err := http.NewRequest("GET", url, nil)
					if err != nil {
						return err
					}
					req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
					req.Header.Add("Accept", "application/json,application/xml")
					res, err := client.Do(req)
					if err != nil {
						return err
					}

					if res.StatusCode != 200 {
						return HandleStatus(res)
					}

					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						return err
					}
					res.Body.Close()

					page := myWorkdayJobsPage{}
					err = json.Unmarshal(body, &page)
					if err != nil {
						return err
					}

					return mwj.readPage(index+lnItems, &page, fn)
				}
			}
		}
	}
	return nil
}

func (mwj *myWorkdayJobs) RetrieveJobs(fn func(job *Job)) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", mwj.url+mwj.home, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json,application/xml")
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return HandleStatus(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	page := myWorkdayJobsPage{}
	err = json.Unmarshal(body, &page)
	if err != nil {
		return err
	}

	return mwj.readPage(0, &page, fn)
}
