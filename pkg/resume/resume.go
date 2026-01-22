// Package resume implements JSON Resume schema support
// See: https://jsonresume.org/schema
package resume

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Resume represents a JSON Resume document
type Resume struct {
	Basics       *Basics       `json:"basics,omitempty"`
	Work         []Work        `json:"work,omitempty"`
	Volunteer    []Volunteer   `json:"volunteer,omitempty"`
	Education    []Education   `json:"education,omitempty"`
	Awards       []Award       `json:"awards,omitempty"`
	Certificates []Certificate `json:"certificates,omitempty"`
	Publications []Publication `json:"publications,omitempty"`
	Skills       []Skill       `json:"skills,omitempty"`
	Languages    []Language    `json:"languages,omitempty"`
	Interests    []Interest    `json:"interests,omitempty"`
	References   []Reference   `json:"references,omitempty"`
	Projects     []Project     `json:"projects,omitempty"`
	Meta         *Meta         `json:"meta,omitempty"`
}

// Basics contains basic personal information
type Basics struct {
	Name     string     `json:"name,omitempty"`
	Label    string     `json:"label,omitempty"`
	Image    string     `json:"image,omitempty"`
	Email    string     `json:"email,omitempty"`
	Phone    string     `json:"phone,omitempty"`
	URL      string     `json:"url,omitempty"`
	Summary  string     `json:"summary,omitempty"`
	Location *Location  `json:"location,omitempty"`
	Profiles []Profile  `json:"profiles,omitempty"`
}

// Location represents a physical address
type Location struct {
	Address     string `json:"address,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	City        string `json:"city,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Region      string `json:"region,omitempty"`
}

// Profile represents a social media or web profile
type Profile struct {
	Network  string `json:"network,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
}

// Work represents a work experience entry
type Work struct {
	Name        string   `json:"name,omitempty"`
	Position    string   `json:"position,omitempty"`
	URL         string   `json:"url,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	Summary     string   `json:"summary,omitempty"`
	Highlights  []string `json:"highlights,omitempty"`
	Location    string   `json:"location,omitempty"`
}

// Volunteer represents a volunteer experience
type Volunteer struct {
	Organization string   `json:"organization,omitempty"`
	Position     string   `json:"position,omitempty"`
	URL          string   `json:"url,omitempty"`
	StartDate    string   `json:"startDate,omitempty"`
	EndDate      string   `json:"endDate,omitempty"`
	Summary      string   `json:"summary,omitempty"`
	Highlights   []string `json:"highlights,omitempty"`
}

// Education represents an education entry
type Education struct {
	Institution string   `json:"institution,omitempty"`
	URL         string   `json:"url,omitempty"`
	Area        string   `json:"area,omitempty"`
	StudyType   string   `json:"studyType,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	Score       string   `json:"score,omitempty"`
	Courses     []string `json:"courses,omitempty"`
}

// Award represents an award or recognition
type Award struct {
	Title   string `json:"title,omitempty"`
	Date    string `json:"date,omitempty"`
	Awarder string `json:"awarder,omitempty"`
	Summary string `json:"summary,omitempty"`
}

// Certificate represents a certification
type Certificate struct {
	Name   string `json:"name,omitempty"`
	Date   string `json:"date,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	URL    string `json:"url,omitempty"`
}

// Publication represents a publication
type Publication struct {
	Name        string `json:"name,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	URL         string `json:"url,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

// Skill represents a skill category
type Skill struct {
	Name     string   `json:"name,omitempty"`
	Level    string   `json:"level,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

// Language represents language proficiency
type Language struct {
	Language string `json:"language,omitempty"`
	Fluency  string `json:"fluency,omitempty"`
}

// Interest represents a personal interest
type Interest struct {
	Name     string   `json:"name,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

// Reference represents a professional reference
type Reference struct {
	Name      string `json:"name,omitempty"`
	Reference string `json:"reference,omitempty"`
}

// Project represents a project
type Project struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Highlights  []string `json:"highlights,omitempty"`
	Keywords    []string `json:"keywords,omitempty"`
	StartDate   string   `json:"startDate,omitempty"`
	EndDate     string   `json:"endDate,omitempty"`
	URL         string   `json:"url,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Entity      string   `json:"entity,omitempty"`
	Type        string   `json:"type,omitempty"`
}

// Meta contains metadata about the resume
type Meta struct {
	Canonical  string `json:"canonical,omitempty"`
	Version    string `json:"version,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
}

// LoadFromFile loads a JSON Resume from a file
func LoadFromFile(path string) (*Resume, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read resume file: %w", err)
	}

	return Parse(data)
}

// Parse parses JSON Resume data
func Parse(data []byte) (*Resume, error) {
	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return nil, fmt.Errorf("failed to parse resume JSON: %w", err)
	}

	return &resume, nil
}

// Validate checks if the resume has required fields
func (r *Resume) Validate() []string {
	var errors []string

	if r.Basics == nil {
		errors = append(errors, "basics section is required")
	} else {
		if r.Basics.Name == "" {
			errors = append(errors, "basics.name is required")
		}
		if r.Basics.Email == "" {
			errors = append(errors, "basics.email is required")
		}
	}

	return errors
}

// ToJSON converts the resume to JSON
func (r *Resume) ToJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// GetSkillKeywords returns all skill keywords as a flat list
func (r *Resume) GetSkillKeywords() []string {
	var keywords []string
	for _, skill := range r.Skills {
		keywords = append(keywords, skill.Keywords...)
	}
	return keywords
}

// GetExperienceYears calculates approximate years of experience
func (r *Resume) GetExperienceYears() int {
	// Simple approximation based on work entries
	return len(r.Work)
}

// MatchesJob checks if the resume skills match job requirements
func (r *Resume) MatchesJob(jobTitle, jobDescription string) float64 {
	if r == nil {
		return 0
	}

	jobTitle = strings.ToLower(jobTitle)
	jobDescription = strings.ToLower(jobDescription)

	var matchScore float64
	var totalKeywords int

	// Check skills match
	for _, skill := range r.Skills {
		for _, keyword := range skill.Keywords {
			totalKeywords++
			kw := strings.ToLower(keyword)
			if strings.Contains(jobTitle, kw) || strings.Contains(jobDescription, kw) {
				matchScore++
			}
		}
	}

	// Check if current/past positions are relevant
	for _, work := range r.Work {
		position := strings.ToLower(work.Position)
		if strings.Contains(jobTitle, position) || strings.Contains(position, jobTitle) {
			matchScore += 2
		}
	}

	if totalKeywords == 0 {
		return 0
	}

	return matchScore / float64(totalKeywords)
}

// Summary returns a brief text summary of the resume
func (r *Resume) Summary() string {
	if r.Basics == nil {
		return "Empty resume"
	}

	var parts []string
	parts = append(parts, r.Basics.Name)

	if r.Basics.Label != "" {
		parts = append(parts, r.Basics.Label)
	}

	if len(r.Work) > 0 {
		parts = append(parts, fmt.Sprintf("%d work experiences", len(r.Work)))
	}

	if len(r.Skills) > 0 {
		parts = append(parts, fmt.Sprintf("%d skill areas", len(r.Skills)))
	}

	return strings.Join(parts, " | ")
}
