package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/k3a/html2text"
)

type greenhouse struct {
	company string
	slug    string
}

// NewGreenhouse returns a new provider
func NewGreenhouse() Provider {
	return &greenhouse{
		"Greenhouse",
		"greenhouse",
	}
}

type greenhouseSearch struct {
	Jobs []struct {
		AbsoluteURL string `json:"absolute_url"`
		Location    struct {
			Name string `json:"name"`
		} `json:"location"`
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"jobs"`
}

type greenhouseJob struct {
	Content string `json:"content"`
}

func (greenhouse *greenhouse) RetrieveJob(job *Job, jobID int, fn func(job *Job)) error {
	res, err := http.Get("https://api.greenhouse.io/v1/boards/" +
		greenhouse.slug + "/embed/job/?id=" + strconv.Itoa(jobID))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("status code error:" + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	search := greenhouseJob{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	job.Desc = html2text.HTML2Text(
		strings.TrimSpace(html.UnescapeString(search.Content)))
	fn(job)
	return nil
}

func (greenhouse *greenhouse) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://api.greenhouse.io/v1/boards/" + greenhouse.slug + "/jobs")
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
	search := greenhouseSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, job := range search.Jobs {
		j := &Job{
			Title:    job.Title,
			Company:  greenhouse.company,
			Location: job.Location.Name,
			Type:     string(FullTime),
			Link:     job.AbsoluteURL,
		}

		err := greenhouse.RetrieveJob(j, job.ID, fn)
		if err != nil {
			return err
		}
	}

	res.Body.Close()
	return nil
}

func (g *greenhouse) ApplyToJob(jobURL string, req *ApplicationRequest) (*ApplicationResult, error) {
	jobID, err := g.extractJobID(jobURL)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("first_name", req.FirstName)
	writer.WriteField("last_name", req.LastName)
	writer.WriteField("email", req.Email)

	if req.Phone != "" {
		writer.WriteField("phone", req.Phone)
	}
	if req.LinkedIn != "" {
		writer.WriteField("urls[LinkedIn]", req.LinkedIn)
	}
	if req.Website != "" {
		writer.WriteField("urls[Website]", req.Website)
	}
	if req.Location != "" {
		writer.WriteField("location", req.Location)
	}

	if len(req.Resume) > 0 {
		part, err := writer.CreateFormFile("resume", "resume.pdf")
		if err != nil {
			return nil, fmt.Errorf("failed to create resume form field: %w", err)
		}
		part.Write(req.Resume)
	}

	if req.CoverLetter != "" {
		writer.WriteField("cover_letter", req.CoverLetter)
	}

	for key, value := range req.CustomFields {
		writer.WriteField(key, value)
	}

	writer.Close()

	url := fmt.Sprintf("https://boards-api.greenhouse.io/v1/boards/%s/jobs/%s", g.slug, jobID)
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
		ID      int    `json:"id"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	json.Unmarshal(respBody, &result)

	return &ApplicationResult{
		Success:       resp.StatusCode < 300,
		ApplicationID: strconv.Itoa(result.ID),
		Message:       result.Message,
	}, nil
}

func (g *greenhouse) extractJobID(jobURL string) (string, error) {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`/jobs/(\d+)`),
		regexp.MustCompile(`job_id=(\d+)`),
		regexp.MustCompile(`#/job/(\d+)`),
	}

	for _, pattern := range patterns {
		if matches := pattern.FindStringSubmatch(jobURL); len(matches) > 1 {
			return matches[1], nil
		}
	}

	return "", fmt.Errorf("could not extract job ID from URL: %s", jobURL)
}
