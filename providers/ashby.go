package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
)

// Ashby provider for companies using Ashby ATS
// API: https://api.ashbyhq.com/posting-api/job-board/{company}

type ashby struct {
	companyName string
	slug        string
}

type ashbyResponse struct {
	Jobs []struct {
		ID                      string `json:"id"`
		Title                   string `json:"title"`
		Location                string `json:"location"`
		Department              string `json:"department"`
		Team                    string `json:"team"`
		EmploymentType          string `json:"employmentType"`
		IsRemote                bool   `json:"isRemote"`
		DescriptionHTML         string `json:"descriptionHtml"`
		DescriptionPlain        string `json:"descriptionPlain"`
		PublishedAt             string `json:"publishedAt"`
		JobURL                  string `json:"jobUrl"`
		ApplyURL                string `json:"applyUrl"`
		CompensationTierSummary string `json:"compensationTierSummary,omitempty"`
	} `json:"jobs"`
}

func (a *ashby) RetrieveJobs(fn func(job *Job)) error {
	url := fmt.Sprintf("https://api.ashbyhq.com/posting-api/job-board/%s", a.slug)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; JobTracker/1.0)")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("ashby API error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response ashbyResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to parse ashby response: %w", err)
	}

	for _, job := range response.Jobs {
		jobType := job.EmploymentType
		if job.IsRemote {
			jobType = "Remote"
		}
		if jobType == "" {
			jobType = string(FullTime)
		}

		fn(&Job{
			Title:    job.Title,
			Company:  a.companyName,
			Location: job.Location,
			Type:     jobType,
			Desc:     job.DescriptionPlain,
			Link:     job.JobURL,
			Misc: map[string]string{
				"department":   job.Department,
				"team":         job.Team,
				"compensation": job.CompensationTierSummary,
				"apply_url":    job.ApplyURL,
			},
		})
	}

	return nil
}

func (a *ashby) ApplyToJob(jobURL string, req *ApplicationRequest) (*ApplicationResult, error) {
	jobID, err := a.extractJobID(jobURL)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("jobPostingId", jobID)
	writer.WriteField("firstName", req.FirstName)
	writer.WriteField("lastName", req.LastName)
	writer.WriteField("email", req.Email)

	if req.Phone != "" {
		writer.WriteField("phoneNumber", req.Phone)
	}
	if req.LinkedIn != "" {
		writer.WriteField("linkedInUrl", req.LinkedIn)
	}
	if req.Website != "" {
		writer.WriteField("websiteUrl", req.Website)
	}
	if req.Location != "" {
		writer.WriteField("currentLocation", req.Location)
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

	url := fmt.Sprintf("https://api.ashbyhq.com/posting-api/job-board/%s/application", a.slug)
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (compatible; JobTracker/1.0)")

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
		Success       bool   `json:"success"`
		Message       string `json:"message"`
	}
	json.Unmarshal(respBody, &result)

	return &ApplicationResult{
		Success:       result.Success || resp.StatusCode < 300,
		ApplicationID: result.ApplicationID,
		Message:       result.Message,
	}, nil
}

func (a *ashby) extractJobID(jobURL string) (string, error) {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`/([a-f0-9-]{36})(?:\?|$)`),
		regexp.MustCompile(`jobs\.ashbyhq\.com/[^/]+/([a-f0-9-]+)`),
	}

	for _, pattern := range patterns {
		if matches := pattern.FindStringSubmatch(jobURL); len(matches) > 1 {
			return matches[1], nil
		}
	}

	return "", fmt.Errorf("could not extract job ID from URL: %s", jobURL)
}

// Specific Ashby-based company types

type openai struct{ ashby }

// NewOpenAI returns a provider for OpenAI jobs
func NewOpenAI() Provider {
	return &openai{ashby{"OpenAI", "openai"}}
}

type notion struct{ ashby }

// NewNotion returns a provider for Notion jobs
func NewNotion() Provider {
	return &notion{ashby{"Notion", "notion"}}
}

type ramp struct{ ashby }

// NewRamp returns a provider for Ramp jobs
func NewRamp() Provider {
	return &ramp{ashby{"Ramp", "ramp"}}
}

type linear struct{ ashby }

// NewLinear returns a provider for Linear jobs
func NewLinear() Provider {
	return &linear{ashby{"Linear", "linear"}}
}

type deel struct{ ashby }

// NewDeel returns a provider for Deel jobs
func NewDeel() Provider {
	return &deel{ashby{"Deel", "deel"}}
}
