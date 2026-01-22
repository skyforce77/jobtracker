package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/k3a/html2text"
	"github.com/pkg/errors"
)

type lever struct {
	company string
	slug    string
}

// NewLever returns a new provider
func NewLever() Provider {
	return &lever{
		"Lever",
		"lever",
	}
}

type leverSearch []struct {
	Title    string `json:"title"`
	Postings []struct {
		Additional      string `json:"additional"`
		AdditionalPlain string `json:"additionalPlain"`
		Categories      struct {
			Commitment string `json:"commitment"`
			Department string `json:"department"`
			Location   string `json:"location"`
			Team       string `json:"team"`
		} `json:"categories"`
		CreatedAt        int64  `json:"createdAt"`
		DescriptionPlain string `json:"descriptionPlain"`
		Description      string `json:"description"`
		ID               string `json:"id"`
		Lists            []struct {
			Text    string `json:"text"`
			Content string `json:"content"`
		} `json:"lists"`
		Text      string `json:"text"`
		HostedURL string `json:"hostedUrl"`
		ApplyURL  string `json:"applyUrl"`
	} `json:"postings"`
}

func (lever *lever) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://api.lever.co/v0/postings/" + lever.slug + "?group=team&mode=json")
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("status code error:" + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	search := leverSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, category := range search {
		for _, job := range category.Postings {
			fn(&Job{
				Title:    job.Text,
				Company:  lever.company,
				Location: job.Categories.Location,
				Type:     job.Categories.Commitment,
				Desc:     strings.TrimSpace(html2text.HTML2Text(job.Description)),
				Link:     job.HostedURL,
			})
		}
	}

	res.Body.Close()
	return nil
}

func (l *lever) ApplyToJob(jobURL string, req *ApplicationRequest) (*ApplicationResult, error) {
	jobID, err := l.extractJobID(jobURL)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("name", req.FirstName+" "+req.LastName)
	writer.WriteField("email", req.Email)

	if req.Phone != "" {
		writer.WriteField("phone", req.Phone)
	}
	if req.LinkedIn != "" {
		writer.WriteField("urls[LinkedIn]", req.LinkedIn)
	}
	if req.Website != "" {
		writer.WriteField("urls[Portfolio]", req.Website)
	}
	if req.Location != "" {
		writer.WriteField("org", req.Location)
	}

	if len(req.Resume) > 0 {
		part, err := writer.CreateFormFile("resume", "resume.pdf")
		if err != nil {
			return nil, fmt.Errorf("failed to create resume form field: %w", err)
		}
		part.Write(req.Resume)
	}

	for key, value := range req.CustomFields {
		writer.WriteField(key, value)
	}

	writer.Close()

	url := fmt.Sprintf("https://api.lever.co/v0/postings/%s/%s", l.slug, jobID)
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return &ApplicationResult{
			Success: false,
			Error:   fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)),
		}, nil
	}

	var result struct {
		ApplicationID string `json:"applicationId"`
		OK            bool   `json:"ok"`
		Message       string `json:"message"`
	}
	json.Unmarshal(respBody, &result)

	return &ApplicationResult{
		Success:       result.OK || resp.StatusCode < 300,
		ApplicationID: result.ApplicationID,
		Message:       result.Message,
	}, nil
}

func (l *lever) extractJobID(jobURL string) (string, error) {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`/postings/[^/]+/([a-f0-9-]+)`),
		regexp.MustCompile(`jobs\.lever\.co/[^/]+/([a-f0-9-]+)`),
	}

	for _, pattern := range patterns {
		if matches := pattern.FindStringSubmatch(jobURL); len(matches) > 1 {
			return matches[1], nil
		}
	}

	return "", fmt.Errorf("could not extract job ID from URL: %s", jobURL)
}
